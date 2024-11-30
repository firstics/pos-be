package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/pos-be/pkg/api/handler"
	"github.com/pos-be/pkg/driver"
	"github.com/pos-be/pkg/usecase"
)

type ErrorHandler struct {
	log driver.Logger
}

func NewErrorHandler(log driver.Logger) *ErrorHandler {
	return &ErrorHandler{log}
}

func (e *ErrorHandler) Handler() gin.HandlerFunc {
	return func(gctx *gin.Context) {
		gctx.Next()
		if len(gctx.Errors) > 0 {
			gerr := gctx.Errors[0].Unwrap()
			e.log.Error(gerr.Error())
			switch e := gerr.(type) {
			case *handler.ErrorBadRequest:
				gctx.JSON(http.StatusBadRequest, e.Error())
				return

			case *usecase.ErrorBusinessException:
				gctx.JSON(http.StatusBadRequest, e.Error())
				return

			default:
				gctx.JSON(http.StatusInternalServerError, e.Error())
				return
			}
		}
	}
}
