package httppack

import (
	"github.com/stretchr/testify/assert"
	"pparse/mock"
	"testing"
	"time"
)

func Test_Packet(t *testing.T) {
	now, _ := time.Parse(time.DateTime, mock.TimeNow)
	var tests = []struct {
		title           string
		packet          ParsedPacket
		isValidExpected bool
	}{
		{"Packet 1", ParsedPacket{URL: "any.com", Uid: 100, At: now}, true},
		{"Packet 2", ParsedPacket{URL: "", Uid: 100, At: now}, true},
		{"Packet 3", ParsedPacket{URL: "abc.com", Uid: 0, At: now}, false},
	}

	for _, test := range tests {
		t.Run(test.title, func(t *testing.T) {
			assert.Equal(t, test.packet.URL, test.packet.Url())
			assert.Equal(t, test.packet.At, test.packet.Timestamp())
			assert.Equal(t, test.packet.Uid, test.packet.ReqRespUid())
			assert.Equal(t, test.packet.IsValid(), test.isValidExpected)
		})
	}
}
