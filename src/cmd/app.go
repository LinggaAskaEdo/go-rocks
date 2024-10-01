package main

import (
	"flag"

	"github.com/go-co-op/gocron/v2"
	"github.com/jmoiron/sqlx"
	"github.com/redis/go-redis/v9"

	domain "github.com/linggaaskaedo/go-rocks/src/business/domain"
	usecase "github.com/linggaaskaedo/go-rocks/src/business/usecase"
	rabbithandler "github.com/linggaaskaedo/go-rocks/src/handler/pubsub/rabbitmq"
	resthandler "github.com/linggaaskaedo/go-rocks/src/handler/rest"
	schedulerhandler "github.com/linggaaskaedo/go-rocks/src/handler/scheduler"
	libauth "github.com/linggaaskaedo/go-rocks/stdlib/auth"
	libgin "github.com/linggaaskaedo/go-rocks/stdlib/gin"
	libgrace "github.com/linggaaskaedo/go-rocks/stdlib/grace"
	libjet "github.com/linggaaskaedo/go-rocks/stdlib/jet"
	libkeycloak "github.com/linggaaskaedo/go-rocks/stdlib/keycloak"
	liblog "github.com/linggaaskaedo/go-rocks/stdlib/logger"
	libmiddleware "github.com/linggaaskaedo/go-rocks/stdlib/middleware"
	libparser "github.com/linggaaskaedo/go-rocks/stdlib/parser"
	preference "github.com/linggaaskaedo/go-rocks/stdlib/preference"
	librabbitmq "github.com/linggaaskaedo/go-rocks/stdlib/rabbitmqlib"
	libredis "github.com/linggaaskaedo/go-rocks/stdlib/redis"
	libscheduler "github.com/linggaaskaedo/go-rocks/stdlib/scheduler"
	libhttpserver "github.com/linggaaskaedo/go-rocks/stdlib/server"
	libsql "github.com/linggaaskaedo/go-rocks/stdlib/sql"
)

var (
	minJitter int
	maxJitter int

	dom *domain.Domain
	uc  *usecase.Usecase

	redisClient0 *redis.Client
	redisClient1 *redis.Client
	redisClient2 *redis.Client

	sqlClient0 *sqlx.DB

	scheduler gocron.Scheduler

	app libgrace.App
)

func init() {
	flag.IntVar(&minJitter, "minSleep", DefaultMinJitter, "min. sleep duration during app initialization")
	flag.IntVar(&maxJitter, "maxSleep", DefaultMaxJitter, "max. sleep duration during app initialization")
	flag.Parse()

	// Add sleep with Jitter to drag the the initialization time among instances
	sleepWithJitter(minJitter, maxJitter)

	// Config Initialization
	conf, err := InitConfig()
	if err != nil {
		panic(err)
	}

	// Logger Initialization
	log := liblog.Init(conf.Log)

	// Redis Initialization
	redisClient0 = libredis.Init(log, conf.Redis, preference.REDIS_APPS)
	redisClient1 = libredis.Init(log, conf.Redis, preference.REDIS_AUTH)
	redisClient2 = libredis.Init(log, conf.Redis, preference.REDIS_LIMITER)

	// SQL Initialization
	sqlClient0 = libsql.Init(log, conf.MySQL)

	// Jet Initialization
	libjet.Init(log, conf.Jet)

	// Auth Initialization
	auth := libauth.Init(log, conf.Auth, redisClient1)

	// Keycloak Initialization
	keycloak := libkeycloak.Init(log, conf.Keycloak)

	// Middleware Initialization
	middleware := libmiddleware.Init(log, conf.Middleware, auth, keycloak, redisClient2)

	// HTTP Gin Initialization
	httpGin := libgin.Init(log, conf.Gin, middleware)

	// Parser Initialization
	parser := libparser.Init(log, conf.Parser)

	// Domain Initialization
	dom = domain.Init(redisClient0, sqlClient0)

	// Usecase Initialization
	uc = usecase.Init(redisClient0, sqlClient0, dom)

	// REST Handler Initialization
	resthandler.Init(httpGin, auth, middleware, parser, keycloak, uc)

	// HTTP Server Initialization
	httpServer := libhttpserver.Init(log, conf.Server, httpGin)

	// Scheduler
	scheduler = libscheduler.Init(log)
	schedulerhandler.Init(log, conf.Scheduler, scheduler)

	// PubSub
	rabbitChannel := librabbitmq.Init(log, conf.RabbitMQ)
	rabbithandler.Init(log, rabbitChannel)

	// App Initialization
	app = libgrace.Init(log, httpServer)
}

// @title			Swagger Go-Rocks
// @version		1.0
// @description	Go-Rocks is a tech playground with Go, Gin
// @contact.name	Lingga Edo
// @contact.url	http://www.github.com/linggaaskaedo
// @contact.email	lemp.otis@gmail.com
func main() {
	defer func() {
		if scheduler != nil {
			scheduler.Shutdown()
		}

		if redisClient0 != nil {
			redisClient0.Close()
		}

		if redisClient1 != nil {
			redisClient1.Close()
		}

		if redisClient2 != nil {
			redisClient2.Close()
		}

		if sqlClient0 != nil {
			sqlClient0.Close()
		}
	}()

	app.Serve()
}
