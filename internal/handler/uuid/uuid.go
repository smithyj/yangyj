package uuid

import (
	"github.com/gin-gonic/gin"
	"github.com/satori/go.uuid"
	"net/http"
	"yangyj/internal/handler"
	"yangyj/pkg/e"
)

type Handler struct{
	handler.Handler
}

func (handler *Handler) UUID() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		u := uuid.NewV4()
		s := u.String()

		handler.JSON(ctx, http.StatusOK, gin.H{
			"code": e.SUCCESS,
			"data": gin.H{
				"uuid": s,
			},
		})
	}
}
