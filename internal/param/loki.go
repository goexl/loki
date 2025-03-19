package param

import (
	"github.com/goexl/http"
	"github.com/goexl/loki/internal/config"
)

type Loki struct {
	Url      string
	Labels   map[string]string
	Batch    *config.Batch
	Username string
	Password string
	Tenant   string
	Http     *http.Client
}

func NewLoki() *Loki {
	return &Loki{
		Labels: make(map[string]string),
		Batch:  config.NewBatch(),
		Http:   http.New().Build(),
	}
}
