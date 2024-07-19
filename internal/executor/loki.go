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

	lokiConfig.Labels = params.Labels
	lokiConfig.Username = params.Username
	lokiConfig.Password = params.Password
	lokiConfig.Tenant = params.Tenant
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
	for _, field := range fields {
		if "" == field.Key() || nil == field.Value() {
			continue
		}

		switch value := field.Value().(type) {
		case bool:
			parsed = append(parsed, zap.Bool(field.Key(), value))
		case *bool:
			parsed = append(parsed, zap.Boolp(field.Key(), value))
		case []bool:
			parsed = append(parsed, zap.Bools(field.Key(), value))
		case *[]bool:
			parsed = append(parsed, zap.Bools(field.Key(), *value))
		case int8:
			parsed = append(parsed, zap.Int8(field.Key(), value))
		case *int8:
			parsed = append(parsed, zap.Int8p(field.Key(), value))
		case int:
			parsed = append(parsed, zap.Int(field.Key(), value))
		case *int:
			parsed = append(parsed, zap.Intp(field.Key(), value))
		case []int:
			parsed = append(parsed, zap.Ints(field.Key(), value))
		case *[]int:
			parsed = append(parsed, zap.Ints(field.Key(), *value))
		case uint:
			parsed = append(parsed, zap.Uint(field.Key(), value))
		case *uint:
			parsed = append(parsed, zap.Uintp(field.Key(), value))
		case []uint:
			parsed = append(parsed, zap.Uints(field.Key(), value))
		case *[]uint:
			parsed = append(parsed, zap.Uints(field.Key(), *value))
		case time.Duration:
			parsed = append(parsed, zap.Duration(field.Key(), value))
		case *time.Duration:
			parsed = append(parsed, zap.Durationp(field.Key(), value))
		case int64:
			parsed = append(parsed, zap.Int64(field.Key(), value))
		case *int64:
			parsed = append(parsed, zap.Int64p(field.Key(), value))
		case []int64:
			parsed = append(parsed, zap.Int64s(field.Key(), value))
		case *[]int64:
			parsed = append(parsed, zap.Int64s(field.Key(), *value))
		case float32:
			parsed = append(parsed, zap.Float32(field.Key(), value))
		case *float32:
			parsed = append(parsed, zap.Float32p(field.Key(), value))
		case float64:
			parsed = append(parsed, zap.Float64(field.Key(), value))
		case *float64:
			parsed = append(parsed, zap.Float64p(field.Key(), value))
		case []float64:
			parsed = append(parsed, zap.Float64s(field.Key(), value))
		case *[]float64:
			parsed = append(parsed, zap.Float64s(field.Key(), *value))
		case *string:
			parsed = append(parsed, zap.Stringp(field.Key(), value))
		case []string:
			parsed = append(parsed, zap.Strings(field.Key(), value))
		case *[]string:
			parsed = append(parsed, zap.Strings(field.Key(), *value))
		case time.Time:
			parsed = append(parsed, zap.Time(field.Key(), value))
		case *time.Time:
			parsed = append(parsed, zap.Timep(field.Key(), value))
		case []time.Time:
			parsed = append(parsed, zap.Times(field.Key(), value))
		case []time.Duration:
			parsed = append(parsed, zap.Durations(field.Key(), value))
		case json.Marshaler, []json.Marshaler:
			// 一定要放在 fmt.Stringer 前面，保证优先使用 json 作为序列化器
			parsed = append(parsed, zap.Reflect(field.Key(), field.Value()))
		case fmt.Stringer:
			parsed = append(parsed, zap.Stringer(field.Key(), value))
		case []fmt.Stringer:
			parsed = append(parsed, zap.Stringers(field.Key(), value))
		case error:
			parsed = append(parsed, zap.Error(value))
		default:
			parsed = append(parsed, zap.Any(field.Key(), field.Value()))
		}
	}

	return
}
