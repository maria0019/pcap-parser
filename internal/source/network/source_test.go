package network

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

const netInterface = "wlp0s20f3"

func Test_Source(t *testing.T) {
	t.Run("Init network source instance", func(t *testing.T) {
		src := NewSource(netInterface)

		assert.Equal(t, src.Type(), TypeNetwork)
		assert.Equal(t, src.Path(), netInterface)
	})

	t.Run("No rights to read from interface", func(t *testing.T) {
		src := NewSource(netInterface)
		_, err := src.Packets()
		assert.Error(t, err)
	})
}
