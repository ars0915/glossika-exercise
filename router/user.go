package router

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/ars0915/glossika-exercise/usecase"
	"github.com/ars0915/glossika-exercise/util/cGin"
)

type registerBody struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,password"`
}

func (rH *HttpHandler) registerHandler(c *gin.Context) {
	ctx := cGin.NewContext(c)

	var body registerBody
	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.WithError(err).Response(http.StatusBadRequest, "Invalid Json")
		return
	}

	if err := rH.h.Register(ctx, usecase.RegisterParam{
		Email:    body.Email,
		Password: body.Password,
	}); err != nil {
		ctx.WithError(err).Response(http.StatusInternalServerError, "Register Failed")
		return
	}

	ctx.Response(http.StatusOK, "")
}

type verifyUserBody struct {
	Email            string `json:"email" binding:"required,email"`
	VerificationCode string `json:"verificationCode" binding:"required,len=6"`
}

func (rH *HttpHandler) verifyUserHandler(c *gin.Context) {
	ctx := cGin.NewContext(c)

	var body verifyUserBody
	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.WithError(err).Response(http.StatusBadRequest, "Invalid Json")
		return
	}

	if err := rH.h.Verify(ctx, usecase.VerifyUserParam{
		Email:            body.Email,
		VerificationCode: body.VerificationCode,
	}); err != nil {
		ctx.WithError(err).Response(http.StatusInternalServerError, "Verify Failed")
		return
	}

	ctx.Response(http.StatusOK, "")
}

type loginBody struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,password"`
}

func (rH *HttpHandler) loginHandler(c *gin.Context) {
	ctx := cGin.NewContext(c)

	var body loginBody
	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.WithError(err).Response(http.StatusBadRequest, "Invalid Json")
		return
	}

	token, err := rH.h.Login(ctx, usecase.LoginParam{
		Email:    body.Email,
		Password: body.Password,
	})
	if err != nil {
		ctx.WithError(err).Response(http.StatusInternalServerError, "Login Failed")
		return
	}
	ctx.Response(http.StatusOK, token)
}
