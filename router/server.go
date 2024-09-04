package router

import (
	"context"
	"net/http"
	"strings"
	"time"

	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"

	"github.com/ars0915/glossika-exercise/config"
	"github.com/ars0915/glossika-exercise/constant"
	"github.com/ars0915/glossika-exercise/util/cGin"
)

func (rH Handler) RunServer(ctx context.Context) (err error) {
	var (
		httpSrv = &http.Server{
			Addr:    ":" + config.Conf.Core.Port,
			Handler: rH.http.routerEngine(),
		}
		errCh = make(chan error)
	)

	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		if err := v.RegisterValidation("password", cGin.ValidatePassword); err != nil {
			errCh <- errors.Wrap(err, "register validation")
		}
	}

	// http server
	go func() {
		cGin.SetResponseCodePrefix(constant.ResponseCodePrefix)

		logrus.Info("HTTP server is running on " + config.Conf.Core.Port + " port.")
		if err := httpSrv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			errCh <- errors.Wrap(err, "listen and serve http")
		}
	}()

	shutdown := func(httpSrv *http.Server) {
		logrus.Warning("Gracefully Shutdown Server ...")

		var (
			finish      []interface{}
			finishCount int
			finishCh    = make(chan interface{})
		)

		timeout, timeoutCancel := context.WithTimeout(context.Background(), 5*time.Minute)
		defer timeoutCancel()

		// Shutdown httpSrv
		finishCount++
		go func() {
			var err error
			defer func() {
				if err != nil {
					finishCh <- err
				} else {
					finishCh <- struct{}{}
				}
			}()
			if err = httpSrv.Shutdown(timeout); err != nil {
				err = errors.Wrap(err, "Shutdown httpSrv")
			}
		}()

		for {
			select {
			case f := <-finishCh:
				if err, ok := f.(error); ok {
					logrus.Error(err)
				}
				if finish = append(finish, f); len(finish) == finishCount {
					return
				}
			case <-timeout.Done():
				logrus.Error("Gracefully Shutdown Timeout.")
				return
			}
		}
	}

	select {
	case err := <-errCh:
		return err
	case <-ctx.Done():
		shutdown(httpSrv)
		return nil
	}
}

func (h HttpHandler) routerEngine() *gin.Engine {
	// set server mode
	gin.SetMode(config.Conf.Core.Mode)

	r := gin.New()
	r.RedirectTrailingSlash = false

	if strings.EqualFold(config.Conf.Core.Mode, "DEBUG") {
		pprof.Register(r)
	}

	r.Use(gin.Recovery())

	r.GET("_health/", func(ctx *gin.Context) {
		ctx.AbortWithStatus(http.StatusOK)
	})

	r.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"text": "Welcome to API server.",
		})
	})

	routers := h.getRouter()
	for _, route := range routers {
		handlerFuncs := []gin.HandlerFunc{}
		if route.requireAuth {
			handlerFuncs = append(handlerFuncs, jwtCheck(h).GinFunc())
		}
		handlerFuncs = append(handlerFuncs, route.worker)

		r.Handle(route.method, route.endpoint, handlerFuncs...)
	}

	return r
}
