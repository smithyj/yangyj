package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"yangyj/pkg/e"
)

func Recovery() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				status := http.StatusInternalServerError
				code := e.ERROR
				msg := e.I18NMsg(ctx.GetHeader("Accept-Language"), code)
				data := gin.H{}
				switch err.(type) {
				case string:
					msg = err.(string)
				case int:
					code = err.(int)
					msg = e.I18NMsg(ctx.GetHeader("Accept-Language"), code)
				case error:
					msg = err.(error).Error()
				case *e.Error:
					eErr := err.(*e.Error)
					if eErr.Status != 0 {
						status = eErr.Status
					}
					if eErr.Code != 0 {
						code = eErr.Code
					}
					if eErr.Msg != "" {
						msg = eErr.Msg
					}
					if len(eErr.Data) > 0 {
						data = eErr.Data
					}
				case e.Error:
					eErr := err.(e.Error)
					if eErr.Status != 0 {
						status = eErr.Status
					}
					if eErr.Code != 0 {
						code = eErr.Code
					}
					if eErr.Msg != "" {
						msg = eErr.Msg
					}
					if len(eErr.Data) > 0 {
						data = eErr.Data
					}
				}
				obj := gin.H{
					"code": code,
				}
				if msg != "" {
					obj["msg"] = msg
				}
				obj["data"] = data
				ctx.JSON(status, obj)
			}
		}()
		ctx.Next()
	}
}
