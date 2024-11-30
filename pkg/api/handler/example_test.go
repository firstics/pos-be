package handler_test

import (
	"encoding/json"
	"errors"
	"net/http"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"

	mocks "github.com/pos-be/mocks/usecase"
	"github.com/pos-be/pkg/api/handler"
)

type exampleHandlerDependencies struct {
	exampleUsecase *mocks.ExampleUsecase
}

func createExampleHandler(t *testing.T) (*handler.ExampleHandler, *exampleHandlerDependencies) {
	exampleUsecase := mocks.NewExampleUsecase(t)
	exampleHandler := handler.NewExampleHandler(exampleUsecase)
	deps := &exampleHandlerDependencies{exampleUsecase}

	return exampleHandler, deps
}

func TestExampleHandler(t *testing.T) {

	method := http.MethodGet
	path := "/example"

	t.Run("It should return 500 on invalid request", func(t *testing.T) {
		exampleHandler, _ := createExampleHandler(t)

		resp, err := request(method, path, strings.NewReader("hello"), exampleHandler.GetText)
		require.NoError(t, err)
		assert.Equal(t, http.StatusInternalServerError, resp.Code)
	})

	t.Run("It should return 200 on valid request", func(t *testing.T) {
		exampleHandler, deps := createExampleHandler(t)
		expected := "hello"

		deps.exampleUsecase.EXPECT().GetText(mock.AnythingOfType("string")).Return(expected, nil)

		// Other ways to construct json body is possible
		body := strings.NewReader(`
			{
				"text": "hello"
			}
		`)
		resp, err := request(method, path, body, exampleHandler.GetText)
		require.NoError(t, err)

		var response handler.ExampleResponse
		err = json.Unmarshal(resp.Body.Bytes(), &response)
		require.NoError(t, err)

		assert.Equal(t, http.StatusOK, resp.Code)
		assert.Equal(t, expected, response.Result)
	})

	t.Run("It should return 500 on usecase error", func(t *testing.T) {
		exampleHandler, deps := createExampleHandler(t)

		deps.exampleUsecase.EXPECT().GetText(mock.AnythingOfType("string")).Return("", errors.New("some error"))

		body := strings.NewReader(`
			{
				"text": "hello"
			}
		`)
		resp, err := request(method, path, body, exampleHandler.GetText)
		require.NoError(t, err)
		assert.Equal(t, http.StatusInternalServerError, resp.Code)
	})

}
