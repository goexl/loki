package core

import (
	"github.com/goexl/log"
	"github.com/goexl/loki/internal/executor"
	"github.com/goexl/loki/internal/param"
)

type Factory struct {
	params *param.Loki
}

func NewFactory(params *param.Loki) *Factory {
	return &Factory{
		params: params,
	}
}

func (f *Factory) New() (log.Executor, error) {
	return executor.NewLoki(f.params)
}
