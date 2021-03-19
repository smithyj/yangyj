package captcha

import (
	gocaptcha "github.com/afocus/captcha"
	"github.com/gin-gonic/gin"
	"image/png"
	"net/http"
	"strconv"
	"yangyj/pkg/captcha"
)

type Handler struct{}

func (handler *Handler) Image() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var err error
		var w, h int
		var img *gocaptcha.Image
		uuid := ctx.Param("uuid")
		if w, err = strconv.Atoi(ctx.DefaultQuery("w", "200")); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"code": 10002,
				"message":  "宽度参数错误",
			})
			return
		}
		if h, err = strconv.Atoi(ctx.DefaultQuery("h", "80")); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"code": 10002,
				"message":  "高度参数错误",
			})
			return
		}
		imageCaptcha := captcha.NewImageCaptcha()
		if img, err = imageCaptcha.Create(uuid, w, h); err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"code": 10001,
				"message":  err.Error(),
			})
			return
		}
		if err = png.Encode(ctx.Writer, img); err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"code": 10001,
				"message":  err.Error(),
			})
			return
		}
	}
}

func (handler *Handler) Email() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var err error
		var req = struct {
			Email string `json:"address" binding:"required,email"`
			UUID string `json:"uuid" binding:"required,uuid"`
			Code string `json:"code" binding:"required,len=6"`
		}{}
		if err = ctx.ShouldBindJSON(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"code": 10002,
				"message": err.Error(),
			})
			return
		}
		imageCaptcha := captcha.NewImageCaptcha()
		if ok := imageCaptcha.Verify(req.UUID, req.Code); !ok {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"code": 10003,
				"message": "图片验证码错误",
			})
			return
		}
		emailCaptcha := captcha.NewEmailCaptcha()
		if err = emailCaptcha.Create(req.Email); err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"code": 10001,
				"message": err.Error(),
			})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{
			"code": 0,
			"message": "成功",
		})
	}
}

func (handler *Handler) Phone() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var err error
		var req = struct {
			CountryNo string `json:"country_no" binding:"required"`
			Phone string `json:"phone" binding:"required"`
			UUID string `json:"uuid" binding:"required,uuid"`
			Code string `json:"code" binding:"required,len=6"`
		}{}
		if err = ctx.ShouldBindJSON(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"code": 10002,
				"message": err.Error(),
			})
			return
		}
		imageCaptcha := captcha.NewImageCaptcha()
		if ok := imageCaptcha.Verify(req.UUID, req.Code); !ok {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"code": 10003,
				"message": "图片验证码错误",
			})
			return
		}

		phoneCaptcha := captcha.NewPhoneCaptcha(req.CountryNo)
		if err = phoneCaptcha.Create(req.Phone); err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"code": 10001,
				"message": err.Error(),
			})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{
			"code": 0,
			"message": "成功",
		})
	}
}
