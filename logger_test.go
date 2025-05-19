package loki_test

import (
	"testing"

	"github.com/goexl/loki"
	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	assert.NotNil(t, loki.New().Build())
}
