package handler_test

import (
	"io"
	"net/http"
	"net/http/httptest"

	"github.com/gin-gonic/gin"

	"github.com/firstics/pos-be/pkg/api/middleware"
	"github.com/firstics/pos-be/pkg/driver"
)

func request(method, path string, body io.Reader, handler gin.HandlerFunc) (*httptest.ResponseRecorder, error) {
	gin.SetMode(gin.TestMode)
	router := gin.New()

	router.Use(middleware.NewErrorHandler(driver.NewNoopLogger()).Handler())

	router.Handle(method, path, handler)

	req, err := http.NewRequest(method, path, body)
	if err != nil {
		return nil, err
	}
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	return resp, nil
}
