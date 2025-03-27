package builder

import (
	"time"

	"github.com/goexl/loki/internal/config"
)

type Batch struct {
	params *config.Batch
	from   *Loki
}

func newBatch(params *config.Batch, from *Loki) *Batch {
	return &Batch{
		params: params,
		from:   from,
	}
}

func (b *Batch) Size(size int) (batch *Batch) {
	b.params.Size = size
	batch = b

	return
}

func (b *Batch) Wait(wait time.Duration) (batch *Batch) {
	b.params.Wait = wait
	batch = b

	return
}

func (b *Batch) Build() *Loki {
	return b.from
}
