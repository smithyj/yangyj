package handler

import (
	"bytes"
	"encoding/base64"
	"fmt"
	gocaptcha "github.com/afocus/captcha"
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
	"image/png"
	"net/http"
	"strconv"
	"time"
	"yangyj/pkg/captcha"
	"yangyj/pkg/e"
)

type CaptchaHandler struct {
	Handler
}

func (handler *CaptchaHandler) Image() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var err error
		var w, h int
		var img *gocaptcha.Image

		if w, err = strconv.Atoi(ctx.DefaultQuery("w", "200")); err != nil {
			handler.JSON(ctx, http.StatusBadRequest, gin.H{
				"code": e.PARAMS_INVALID,
			})
			return
		}
		if h, err = strconv.Atoi(ctx.DefaultQuery("h", "80")); err != nil {
			handler.JSON(ctx, http.StatusBadRequest, gin.H{
				"code": e.PARAMS_INVALID,
			})
			return
		}

		u := uuid.NewV4()
		u = uuid.NewV5(u, strconv.Itoa(time.Now().Nanosecond()))
		s := u.String()

		imageCaptcha := captcha.NewImageCaptcha()
		if img, err = imageCaptcha.Create(s, w, h); err != nil {
			handler.JSON(ctx, http.StatusInternalServerError, gin.H{
				"code": e.ERROR,
			})
			return
		}
		buf := bytes.NewBuffer([]byte{})
		if err = png.Encode(buf, img); err != nil {
			handler.JSON(ctx, http.StatusInternalServerError, gin.H{
				"code": e.ERROR,
				"msg":  err.Error(),
			})
			return
		}
		base64Str := fmt.Sprintf("data:image/jpeg;base64,%v", base64.StdEncoding.EncodeToString(buf.Bytes()))

		handler.JSON(ctx, http.StatusOK, gin.H{
			"code": e.SUCCESS,
			"data": gin.H{
				"captchaId": s,
				"base64Png": base64Str,
			},
		})
	}
}

func (handler *CaptchaHandler) Email() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var err error
		var req = struct {
			Email string `json:"email" binding:"required,email"`
			UUID  string `json:"uuid" binding:"required,uuid"`
			Code  string `json:"code" binding:"required,len=6"`
		}{}
		if err = ctx.ShouldBindJSON(&req); err != nil {
			handler.JSON(ctx, http.StatusBadRequest, gin.H{
				"code": e.PARAMS_INVALID,
				"msg":  err.Error(),
			})
			return
		}
		imageCaptcha := captcha.NewImageCaptcha()
		if ok := imageCaptcha.Verify(req.UUID, req.Code); !ok {
			handler.JSON(ctx, http.StatusBadRequest, gin.H{
				"code": e.CAPTCHA_INVALID,
			})
			return
		}
		emailCaptcha := captcha.NewEmailCaptcha()
		if err = emailCaptcha.Create(req.Email); err != nil {
			handler.JSON(ctx, http.StatusInternalServerError, gin.H{
				"code": e.ERROR,
				"msg":  err.Error(),
			})
			return
		}
		handler.JSON(ctx, http.StatusOK, gin.H{
			"code": e.SUCCESS,
		})
	}
}

func (handler *CaptchaHandler) Phone() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var err error
		var req = struct {
			CountryCode string `json:"country_code" binding:"required"`
			Phone       string `json:"phone" binding:"required"`
			UUID        string `json:"uuid" binding:"required,uuid"`
			Code        string `json:"code" binding:"required,len=6"`
		}{}
		if err = ctx.ShouldBindJSON(&req); err != nil {
			handler.JSON(ctx, http.StatusBadRequest, gin.H{
				"code": e.PARAMS_INVALID,
				"msg":  err.Error(),
			})
			return
		}
		imageCaptcha := captcha.NewImageCaptcha()
		if ok := imageCaptcha.Verify(req.UUID, req.Code); !ok {
			handler.JSON(ctx, http.StatusBadRequest, gin.H{
				"code": e.CAPTCHA_INVALID,
			})
			return
		}

		phoneCaptcha := captcha.NewPhoneCaptcha(req.CountryCode)
		if err = phoneCaptcha.Create(req.Phone); err != nil {
			handler.JSON(ctx, http.StatusInternalServerError, gin.H{
				"code": e.ERROR,
				"msg":  err.Error(),
			})
			return
		}
		handler.JSON(ctx, http.StatusOK, gin.H{
			"code": e.SUCCESS,
		})
	}
}
