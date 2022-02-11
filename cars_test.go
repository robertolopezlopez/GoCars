package awesomeProject

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

type (
	mockremoteservice struct {
		mock.Mock
	}
)

func TestIncrementVersion(t *testing.T) {
	t.Run("1 to 2", func(t *testing.T) {
		m := NewCar().IncrementVersion()
		assert.Equal(t, 2, m.Version)
	})
	t.Run("1 to 3", func(t *testing.T) {
		m := NewCar().IncrementVersion()
		assert.NotEqual(t, 3, m.Version)
	})
}
