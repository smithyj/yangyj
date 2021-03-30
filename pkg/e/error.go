package e

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

type Error struct {
	Status  int
	Code    int
	Msg string
	Data    gin.H
}

func (e *Error) Error() string {
	return fmt.Sprintf("%v", e.Code)
}
