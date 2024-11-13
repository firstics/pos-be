package usecase

import (
	"github.com/firstics/pos-be/pkg/driver"
)

type ExampleUsecase interface {
	GetText(text string) (string, error)
}

type exampleUsecase struct {
	log driver.Logger
}

func NewExampleUsecase(log driver.Logger) ExampleUsecase {
	return &exampleUsecase{log}
}

func (eu *exampleUsecase) GetText(text string) (string, error) {
	return text, nil
}
