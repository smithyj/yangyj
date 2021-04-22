package handler

import (
	"github.com/gin-gonic/gin"
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
	engine.GET("/captcha/image", new(CaptchaHandler).Image())
	engine.POST("/captcha/email", new(CaptchaHandler).Email())
	engine.POST("/captcha/phone", new(CaptchaHandler).Phone())
}

func Router() (engine *gin.Engine) {
	engine = new(router).init()
	return
}
