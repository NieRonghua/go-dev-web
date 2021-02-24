package api

import (
	"github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	"github.com/itsjamie/gin-cors"
	"go-dev-web/pkg/api/controller"
	"go-dev-web/pkg/plog"
	"time"
)

func ServeAPI(cfg *Config) error {
	r := gin.New()

	logger := plog.Logger()
	r.Use(ginzap.Ginzap(logger, time.RFC3339, true))
	r.Use(ginzap.RecoveryWithZap(logger, true))

	r.Use(cors.Middleware(cors.Config{
		Origins:         "*",
		Methods:         "GET, PUT, POST, DELETE, OPTIONS",
		RequestHeaders:  "Origin, Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, If-Modified-Since",
		ExposedHeaders:  "",
		MaxAge:          50 * time.Second,
		Credentials:     false,
		ValidateHeaders: false,
	}))
	apiGroup := r.Group("/")
	{
		apiGroup.GET("/", controller.Index)
		apiGroup.POST("/hello", controller.HelloWorld)
	}

	return r.Run(cfg.BindAddr)
}
