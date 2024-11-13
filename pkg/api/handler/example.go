package handler

import (
	"github.com/gin-gonic/gin"

	"github.com/firstics/pos-be/pkg/usecase"
)

type ExampleRequester struct {
	Text string `json:"text"`
}

type ExampleResponse struct {
	Result string `json:"result"`
}

type ExampleHandler struct {
	exampleUsecase usecase.ExampleUsecase
}

func NewExampleHandler(usecase usecase.ExampleUsecase) *ExampleHandler {
	return &ExampleHandler{
		exampleUsecase: usecase,
	}
}

func (eh *ExampleHandler) GetText(c *gin.Context) {
	var req ExampleRequester
	if err := c.ShouldBindJSON(&req); err != nil {
		c.Error(err)
		return
	}

	result, err := eh.exampleUsecase.GetText(req.Text)
	if err != nil {
		c.Error(err)
		return
	}

	var resp ExampleResponse
	resp.Result = result
	c.JSON(200, resp)
}
