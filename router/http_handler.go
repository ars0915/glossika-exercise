package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type appRouter struct {
	method      string
	endpoint    string
	requireAuth bool
	worker      gin.HandlerFunc
}

func (h HttpHandler) getRouter() (routes []appRouter) {
	return []appRouter{
		{http.MethodPost, "/register", false, h.registerHandler},
		{http.MethodPost, "/verify", false, h.verifyUserHandler},
		{http.MethodPost, "/login", false, h.loginHandler},
	}
}
