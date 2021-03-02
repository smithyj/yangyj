package uuid

import (
	"github.com/gin-gonic/gin"
	"github.com/satori/go.uuid"
	"net/http"
)

type Handler struct {}

func (handler *Handler) UUID() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		u := uuid.NewV4()
		ctx.JSON(http.StatusOK, gin.H{
			"code": 0,
			"data": gin.H{
				"uuid": u.String(),
			},
		})
	}
}
