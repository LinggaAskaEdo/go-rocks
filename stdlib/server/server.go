package server

import (
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
)

var once = &sync.Once{}

type httpServer struct {
	logger  log.Logger
	servers []*http.Server
	opt     Options
}

type Options struct {
	Port         int `json:"port"`
	WriteTimeout int `json:"write_timeout"`
	ReadTimeout  int `json:"read_timeout"`
	IdleTimeout  int `json:"idle_timeout"`
}

func Init(log zerolog.Logger, opt Options, gin *gin.Engine) *http.Server {
	var server *http.Server

	once.Do(func() {
		serverPort := fmt.Sprintf(":%d", opt.Port)

		server = &http.Server{
			Addr:         serverPort,
			WriteTimeout: time.Second * time.Duration(opt.WriteTimeout),
			ReadTimeout:  time.Second * time.Duration(opt.ReadTimeout),
			IdleTimeout:  time.Second * time.Duration(opt.IdleTimeout),
			Handler:      gin,
		}
	})

	return server
}
