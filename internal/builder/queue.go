package builder

import (
	"github.com/goexl/loki/internal/config"
)

type Queue struct {
	params *config.Queue
	from   *Loki
}

func newQueue(params *config.Queue, from *Loki) *Queue {
	return &Queue{
		params: params,
		from:   from,
	}
}

func (q *Queue) Name(name string) (queue *Queue) {
	q.params.Name = name
	queue = q

	return
}

func (q *Queue) Directory(directory string) (queue *Queue) {
	q.params.Directory = directory
	queue = q

	return
}

func (q *Queue) Build() *Loki {
	return q.from
}
