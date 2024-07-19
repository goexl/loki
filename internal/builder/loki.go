package builder

import (
	"time"

	"github.com/goexl/http"
	"github.com/goexl/loki/internal/core"
	"github.com/goexl/loki/internal/param"
)

type Loki struct {
	params *param.Loki
}

func NewLoki() *Loki {
	return &Loki{
		params: param.NewLoki(),
	}
}

func (l *Loki) Url(url string) (loki *Loki) {
	l.params.Url = url
	loki = l

	return
}

func (l *Loki) Username(username string) (loki *Loki) {
	l.params.Username = username
	loki = l

	return
}

func (l *Loki) Password(password string) (loki *Loki) {
	l.params.Password = password
	loki = l

	return
}

func (l *Loki) Batch(size int, wait time.Duration) (loki *Loki) {
	l.params.Batch.Size = size
	l.params.Batch.Wait = wait
	loki = l

	return
}

func (l *Loki) Labels(labels map[string]string) (loki *Loki) {
	l.params.Labels = labels
	loki = l

	return
}

func (l *Loki) Label(key string, value string) (loki *Loki) {
	l.params.Labels[key] = value
	loki = l

	return
}

func (l *Loki) Tenant(tenant string) (loki *Loki) {
	l.params.Tenant = tenant
	loki = l

	return
}

func (l *Loki) Http(http *http.Client) (loki *Loki) {
	l.params.Http = http
	loki = l

	return
}

func (l *Loki) Build() *core.Factory {
	return core.NewFactory(l.params)
}
