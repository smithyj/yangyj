package handler

import (
	"github.com/gin-gonic/gin"
	"yangyj/pkg/e"
)

type Handler struct{}

func (h *Handler) JSON(ctx *gin.Context, status int, obj gin.H) {
	if _, ok := obj["msg"]; !ok {
		obj["msg"] = ""
		if v, ok := obj["code"]; ok {
			obj["msg"] = e.I18NMsg(ctx.GetHeader("Accept-Language"), v.(int))
		}
	}
	ctx.JSON(status, obj)
}
