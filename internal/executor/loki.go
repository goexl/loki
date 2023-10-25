package executor

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/goexl/gox"
	"github.com/goexl/loki/internal/executor/internal/config"
	"github.com/goexl/loki/internal/internal/loki"
	"github.com/goexl/loki/internal/param"
	"go.uber.org/zap"
)

type Loki struct {
	zap *zap.Logger
}

func NewLoki(params *param.Loki) (logger *Loki, err error) {
	logger = new(Loki)
	lokiConfig := new(loki.Config)
	lokiConfig.Url = params.Url
	lokiConfig.Batch = params.Batch
	lokiConfig.Http = params.Http

	if 0 != len(params.Labels) {
		lokiConfig.Labels = params.Labels
	}
	if "" != params.Username {
		lokiConfig.Username = params.Username
	}
	if "" != params.Password {
		lokiConfig.Password = params.Password
	}
	pusher := loki.New(context.Background(), lokiConfig)
	logger.zap, err = pusher.Build(config.DefaultZap(), zap.WithCaller(false))

	return
}

func (l *Loki) Debug(msg string, fields ...gox.Field[any]) {
	l.zap.Debug(msg, l.parse(fields...)...)
}

func (l *Loki) Info(msg string, fields ...gox.Field[any]) {
	l.zap.Info(msg, l.parse(fields...)...)
}

func (l *Loki) Warn(msg string, fields ...gox.Field[any]) {
	l.zap.Warn(msg, l.parse(fields...)...)
}

func (l *Loki) Error(msg string, fields ...gox.Field[any]) {
	l.zap.Error(msg, l.parse(fields...)...)
}

func (l *Loki) Panic(msg string, fields ...gox.Field[any]) {
	l.zap.Panic(msg, l.parse(fields...)...)
}

func (l *Loki) Fatal(msg string, fields ...gox.Field[any]) {
	l.zap.Fatal(msg, l.parse(fields...)...)
}

func (l *Loki) Sync() error {
	return l.zap.Sync()
}

func (l *Loki) parse(fields ...gox.Field[any]) (parsed []zap.Field) {
	parsed = make([]zap.Field, 0, len(fields))
	for _, f := range fields {
		if "" == f.Key() || nil == f.Value() {
			continue
		}

		switch value := f.Value().(type) {
		case bool:
			parsed = append(parsed, zap.Bool(f.Key(), value))
		case *bool:
			parsed = append(parsed, zap.Boolp(f.Key(), value))
		case []bool:
			parsed = append(parsed, zap.Bools(f.Key(), value))
		case *[]bool:
			parsed = append(parsed, zap.Bools(f.Key(), *value))
		case int8:
			parsed = append(parsed, zap.Int8(f.Key(), value))
		case *int8:
			parsed = append(parsed, zap.Int8p(f.Key(), value))
		case int:
			parsed = append(parsed, zap.Int(f.Key(), value))
		case *int:
			parsed = append(parsed, zap.Intp(f.Key(), value))
		case []int:
			parsed = append(parsed, zap.Ints(f.Key(), value))
		case *[]int:
			parsed = append(parsed, zap.Ints(f.Key(), *value))
		case uint:
			parsed = append(parsed, zap.Uint(f.Key(), value))
		case *uint:
			parsed = append(parsed, zap.Uintp(f.Key(), value))
		case []uint:
			parsed = append(parsed, zap.Uints(f.Key(), value))
		case *[]uint:
			parsed = append(parsed, zap.Uints(f.Key(), *value))
		case int64:
			parsed = append(parsed, zap.Int64(f.Key(), value))
		case *int64:
			parsed = append(parsed, zap.Int64p(f.Key(), value))
		case []int64:
			parsed = append(parsed, zap.Int64s(f.Key(), value))
		case *[]int64:
			parsed = append(parsed, zap.Int64s(f.Key(), *value))
		case float32:
			parsed = append(parsed, zap.Float32(f.Key(), value))
		case *float32:
			parsed = append(parsed, zap.Float32p(f.Key(), value))
		case float64:
			parsed = append(parsed, zap.Float64(f.Key(), value))
		case *float64:
			parsed = append(parsed, zap.Float64p(f.Key(), value))
		case []float64:
			parsed = append(parsed, zap.Float64s(f.Key(), value))
		case *[]float64:
			parsed = append(parsed, zap.Float64s(f.Key(), *value))
		case *string:
			parsed = append(parsed, zap.Stringp(f.Key(), value))
		case []string:
			parsed = append(parsed, zap.Strings(f.Key(), value))
		case *[]string:
			parsed = append(parsed, zap.Strings(f.Key(), *value))
		case time.Time:
			parsed = append(parsed, zap.Time(f.Key(), value))
		case *time.Time:
			parsed = append(parsed, zap.Timep(f.Key(), value))
		case []time.Time:
			parsed = append(parsed, zap.Times(f.Key(), value))
		case time.Duration:
			parsed = append(parsed, zap.Duration(f.Key(), value))
		case *time.Duration:
			parsed = append(parsed, zap.Durationp(f.Key(), value))
		case []time.Duration:
			parsed = append(parsed, zap.Durations(f.Key(), value))
		case json.Marshaler, []json.Marshaler:
			// 一定要放在 fmt.Stringer 前面，保证优先使用 json 作为序列化器
			parsed = append(parsed, zap.Reflect(f.Key(), f.Value()))
		case fmt.Stringer:
			parsed = append(parsed, zap.Stringer(f.Key(), value))
		case []fmt.Stringer:
			parsed = append(parsed, zap.Stringers(f.Key(), value))
		case error:
			parsed = append(parsed, zap.Error(value))
		default:
			parsed = append(parsed, zap.Any(f.Key(), f.Value()))
		}
	}

	return
}
