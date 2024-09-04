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
