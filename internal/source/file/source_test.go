package file

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

const filePath = "any.pcap"

func Test_Source(t *testing.T) {
	t.Run("Init file source instance", func(t *testing.T) {
		src := NewSource(filePath)

		assert.Equal(t, src.Type(), TypeFile)
		assert.Equal(t, src.Path(), filePath)
	})

	t.Run("No .pcap file found", func(t *testing.T) {
		src := NewSource(filePath)
		_, err := src.Packets()
		assert.Error(t, err)
	})
}
