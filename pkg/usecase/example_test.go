package usecase_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	mocks "github.com/firstics/pos-be/mocks/driver"
	"github.com/firstics/pos-be/pkg/usecase"
)

type exampleDependencies struct {
	mockLog *mocks.Logger
}

func createExampleUsecase(t *testing.T) (usecase.ExampleUsecase, *exampleDependencies) {
	mockLog := mocks.NewLogger(t)
	exampleUsecase := usecase.NewExampleUsecase(mockLog)
	return exampleUsecase, &exampleDependencies{mockLog}
}

func TestExampleGetText(t *testing.T) {
	t.Run("It should return err as nil if get text works properly", func(t *testing.T) {
		exampleUsecase, deps := createExampleUsecase(t)
		text := "Hello world"
		expectedResult := "Hello world"
		res, err := exampleUsecase.GetText(text)

		assert.NoError(t, err)
		assert.Equal(t, expectedResult, res)

		deps.mockLog.AssertExpectations(t)
	})
}
