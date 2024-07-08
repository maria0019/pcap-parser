package source

import (
	"github.com/stretchr/testify/assert"
	"pparse/internal/source/file"
	"pparse/internal/source/network"
	"testing"
)

func Test_Source(t *testing.T) {
	f, err := NewDataSource("any.pcap", "")
	i, ok := f.(file.Source)
	assert.True(t, ok, "Should be file source instance", i)
	assert.NoError(t, err)

	n, err := NewDataSource("", "wlp0s20f3")
	_, ok = n.(network.Source)
	assert.True(t, ok, "Should be network source instance")
	assert.NoError(t, err)

	empty, err := NewDataSource("", "")
	assert.Empty(t, empty, "Do source data should be returned")
	assert.Error(t, err)
}
