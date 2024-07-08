package source

import (
	"github.com/stretchr/testify/assert"
	calculatorMock "pparse/internal/calculator/mock"
	"pparse/internal/source/file"
	"pparse/internal/source/network"
	"testing"
)

func Test_Counter(t *testing.T) {
	calc := calculatorMock.Calculator{}
	timeInterval := 10

	f, err := NewCounter("any.pcap", "", calc, timeInterval)
	_, ok := f.(*file.Counter)
	assert.True(t, ok, "Should be file counter instance")
	assert.NoError(t, err)

	n, err := NewCounter("", "wlp0s20f3", calc, timeInterval)
	_, ok = n.(*network.Counter)
	assert.True(t, ok, "Should be network counter instance")
	assert.NoError(t, err)

	empty, err := NewCounter("", "", calc, timeInterval)
	assert.Empty(t, empty, "Do counter instance should be returned")
	assert.Error(t, err)
}
