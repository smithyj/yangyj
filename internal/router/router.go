package router

import (
	"github.com/gin-gonic/gin"
	"yangyj/internal/handler/captcha"
	"yangyj/internal/handler/uuid"
	"yangyj/internal/middleware"
	"yangyj/pkg/config"
)

type router struct{}

func (r *router) init() (engine *gin.Engine) {
	gin.SetMode(config.Config.Mode)
	engine = gin.New()
	engine.Use(gin.Logger())
	engine.Use(gin.Recovery())
	engine.Use(middleware.Recovery())
	r.inject(engine)
	return
}

func (r *router) inject(engine *gin.Engine) {
	engine.GET("/uuid", new(uuid.Handler).UUID())
	engine.GET("/captcha/image/:uuid", new(captcha.Handler).Image())
	engine.POST("/captcha/email", new(captcha.Handler).Email())
	engine.POST("/captcha/phone", new(captcha.Handler).Phone())
}

func New() (engine *gin.Engine) {
	engine = new(router).init()
	return
}
