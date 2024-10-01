package middleware

import (
	"context"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rs/xid"

	"github.com/linggaaskaedo/go-rocks/stdlib/preference"
)

func (mw *middleware) Handler() gin.HandlerFunc {
	return func(c *gin.Context) {
		path := c.Request.URL.Path

		if !strings.HasPrefix(path, "/swagger/") { // skip logging swagger request
			start := time.Now()

			ctx := c.Request.Context()
			ctx = mw.attachReqID(ctx)
			ctx = mw.attachLogger(ctx)

			raw := c.Request.URL.RawQuery
			if raw != "" {
				path = path + "?" + raw
			}

			mw.log.Info().
				Str("event", "START").
				Str("correlation_id", mw.getRequestID(ctx)).
				Str("method", c.Request.Method).
				Str("url", path).
				Str("user_agent", c.Request.UserAgent()).
				Str("addr", c.Request.Host).
				Send()

			// Process request
			c.Request = c.Request.WithContext(ctx)
			c.Next()

			// Fill the params
			param := gin.LogFormatterParams{}

			param.TimeStamp = time.Now() // Stop timer
			param.Latency = param.TimeStamp.Sub(start)
			if param.Latency > time.Minute {
				param.Latency = param.Latency.Truncate(time.Second)
			}

			param.StatusCode = c.Writer.Status()

			mw.log.Info().
				Str("event", "END").
				Str("correlation_id", mw.getRequestID(ctx)).
				Str("latency", param.Latency.String()).
				Int("status_code", param.StatusCode).
				Send()
		}
	}
}

func (mw *middleware) attachReqID(ctx context.Context) context.Context {
	return context.WithValue(ctx, preference.ContextKeyRequestID, xid.New().String())
}

func (mw *middleware) attachLogger(ctx context.Context) context.Context {
	return mw.log.With().Str(preference.ContextKeyCorrelationID, mw.getRequestID(ctx)).Logger().WithContext(ctx)
}

func (mw *middleware) getRequestID(ctx context.Context) string {
	reqID := ctx.Value(preference.ContextKeyRequestID)

	if ret, ok := reqID.(string); ok {
		return ret
	}

	return ""
}
