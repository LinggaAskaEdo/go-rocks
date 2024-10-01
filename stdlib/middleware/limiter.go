package middleware

import (
	"context"
	"errors"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/linggaaskaedo/go-rocks/stdlib/preference"
)

func (mw *middleware) Limiter(command string, limit int) gin.HandlerFunc {
	t, _ := mw.parseCommand(command)

	return func(c *gin.Context) {
		now := time.Now().Unix()
		clientIp := c.ClientIP()
		deadline := mw.getDeadLine()
		routeDeadline := time.Now().Add(t).Unix()
		routeKey := c.FullPath() + c.Request.Method + clientIp // for single route limit in redis.
		staticKey := clientIp                                  // for global limit search in redis.

		routeLimit := limit
		staticLimit := mw.limit

		keys := []string{routeKey, staticKey}
		args := []interface{}{routeLimit, staticLimit, routeDeadline, now}

		// mean global limit should be reset.
		if now > deadline {
			mw.updateDeadLine()
			_, err := mw.rdb.EvalSha(context.Background(), mw.getSHAScript("reset"), keys, routeDeadline).Result()
			if err != nil {
				c.JSON(http.StatusInternalServerError, err)
				c.Abort()
			}

			c.Header("X-RateLimit-Limit-global", strconv.Itoa(staticLimit))
			c.Header("X-RateLimit-Remaining-global", strconv.Itoa(staticLimit-1))
			c.Header("X-RateLimit-Reset-global", mw.getDeadLineWithString())
			c.Header("X-RateLimit-Limit-route", strconv.Itoa(limit))
			c.Header("X-RateLimit-Remaining-route", strconv.Itoa(limit-1))
			c.Header("X-RateLimit-Reset-route", time.Unix(routeDeadline, 0).Format(TimeFormat))
			c.Next()
			return
		}

		results, err := mw.rdb.EvalSha(context.Background(), mw.getSHAScript("normal"), keys, args).Result()
		if err != nil {
			c.JSON(http.StatusInternalServerError, err)
			c.Abort()
		}

		result := results.([]interface{})
		routeRemaining := result[0].(int64)
		staticRemaining := result[1].(int64)
		routedeadline := time.Unix(result[2].(int64), 0).Format(TimeFormat)

		if routeRemaining == -1 {
			c.JSON(http.StatusTooManyRequests, routedeadline)
			c.Header("X-RateLimit-Reset-single", routedeadline)
			c.Abort()
			return
		}

		if staticRemaining == -1 {
			c.JSON(http.StatusTooManyRequests, mw.getDeadLineWithString())
			c.Header("X-RateLimit-Reset-global", mw.getDeadLineWithString())
			c.Abort()
			return
		}

		c.Header("X-RateLimit-Limit-global", strconv.Itoa(staticLimit))
		c.Header("X-RateLimit-Remaining-global", strconv.FormatInt(staticRemaining, 10))
		c.Header("X-RateLimit-Reset-global", mw.getDeadLineWithString())
		c.Header("X-RateLimit-Limit-route", strconv.Itoa(routeLimit))
		c.Header("X-RateLimit-Remaining-route", strconv.FormatInt(routeRemaining, 10))
		c.Header("X-RateLimit-Reset-route", routedeadline)
		c.Next()
	}
}

func (mw *middleware) parseCommand(command string) (time.Duration, error) {
	var period time.Duration

	values := strings.Split(command, "-")
	if len(values) != 2 {
		return period, errors.New(preference.FormatError)
	}

	unit, err := strconv.Atoi(values[0])
	if err != nil {
		return period, errors.New(preference.FormatError)
	}

	if unit <= 0 {
		return period, errors.New(preference.CommandError)
	}

	if t, ok := timeDict[strings.ToUpper(values[1])]; ok {
		return time.Duration(unit) * t, nil
	} else {
		return period, errors.New(preference.FormatError)
	}
}

func (mw *middleware) getDeadLine() int64 {
	return mw.deadline
}

func (mw *middleware) updateDeadLine() {
	mw.deadline = time.Now().Add(mw.period).Unix()
}

func (mw *middleware) getSHAScript(index string) string {
	return mw.shaScript[index]
}

func (mw *middleware) getDeadLineWithString() string {
	return time.Unix(mw.deadline, 0).Format(TimeFormat)
}
