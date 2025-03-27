package config

type Queue struct {
	Name      string
	Directory string
}

func NewQueue() *Queue {
	return &Queue{
		Name:      "loki",
		Directory: ".",
	}
}
