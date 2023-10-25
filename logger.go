package loki

import (
	"github.com/goexl/log"
	"github.com/goexl/loki/internal/builder"
)

func New() *builder.Loki {
	return builder.NewLoki()
}

// Logger 日志接口
type Logger = log.Logger
