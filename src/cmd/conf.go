package main

import (
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"

	handlerscheduler "github.com/linggaaskaedo/go-rocks/src/handler/scheduler"
	libauth "github.com/linggaaskaedo/go-rocks/stdlib/auth"
	libgin "github.com/linggaaskaedo/go-rocks/stdlib/gin"
	libjet "github.com/linggaaskaedo/go-rocks/stdlib/jet"
	libkeycloak "github.com/linggaaskaedo/go-rocks/stdlib/keycloak"
	liblog "github.com/linggaaskaedo/go-rocks/stdlib/logger"
	libmiddleware "github.com/linggaaskaedo/go-rocks/stdlib/middleware"
	libparser "github.com/linggaaskaedo/go-rocks/stdlib/parser"
	librabbit "github.com/linggaaskaedo/go-rocks/stdlib/rabbitmqlib"
	libredis "github.com/linggaaskaedo/go-rocks/stdlib/redis"
	libhttpserver "github.com/linggaaskaedo/go-rocks/stdlib/server"
	libsql "github.com/linggaaskaedo/go-rocks/stdlib/sql"
)

type Config struct {
	Log        liblog.Options
	Redis      libredis.Options
	MySQL      libsql.Options
	RabbitMQ   librabbit.Options
	Jet        libjet.Options
	Auth       libauth.Options
	Middleware libmiddleware.Options
	Keycloak   libkeycloak.Options
	Gin        libgin.Options
	Parser     libparser.Options
	Server     libhttpserver.Options
	Scheduler  handlerscheduler.Options
}

func InitConfig() (config Config, err error) {
	cfg := &Config{}

	err = godotenv.Load()
	if err != nil {
		return
	}

	err = envconfig.Process("", cfg)
	if err != nil {
		return
	}

	config = *cfg

	return
}
