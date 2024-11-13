package di

import (
	"github.com/google/wire"

	"github.com/firstics/pos-be/pkg/driver"
)

var LogSet = wire.NewSet(
	driver.NewLogger,
)
