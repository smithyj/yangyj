package router

import (
	"github.com/gin-gonic/gin"
	"yangyj/internal/handler/uuid"
)

type router struct {}

func (r *router) init() (engine *gin.Engine) {
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