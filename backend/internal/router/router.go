package router

import (
	"github.com/gin-gonic/gin"
	"yangyj/backend/internal/handler/uuid"
	"yangyj/backend/pkg/config"
)

type router struct {}

func (r *router) init() (engine *gin.Engine) {
	gin.SetMode(config.Config.Mode)
	engine = gin.New()
	engine.Use(gin.Logger())
	engine.Use(gin.Recovery())
	r.inject(engine)
	return
}

func (r *router) inject(engine *gin.Engine) {
	engine.GET("/uuid", new(uuid.Handler).UUID())
}

func New() (engine *gin.Engine) {
	engine = new(router).init()
	return
}