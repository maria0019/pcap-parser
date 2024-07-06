package httppack

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func Test_Calculator(t *testing.T) {
	now, _ := time.Parse(time.DateTime, "2020-01-01 00:00:00")
	var tests = []struct {
		title    string
		at       time.Time
		packets  []ParsedPacket
		expected Stats
	}{
		{"Interval 1",
			now.Add(time.Duration(5) * time.Second),
			[]ParsedPacket{
				{URL: "any.com", Uid: 100, At: now.Add(2 * time.Millisecond)},
				{URL: "any.com", Uid: 100, At: now.Add(4 * time.Millisecond)}, // 2 Millisecond resp time
				{URL: "any.com", Uid: 101, At: now.Add(6 * time.Millisecond)},
				{URL: "any.com", Uid: 101, At: now.Add(16 * time.Millisecond)}, // 10 Millisecond resp time
				{URL: "abc.com", Uid: 102, At: now.Add(10 * time.Millisecond)},
				{URL: "abc.com", Uid: 102, At: now.Add(18 * time.Millisecond)}, // 8 Millisecond resp time
			},
			Stats{
				At:                "2020-01-01 00:00:05",
				AvgResponseTimeMs: 6,
				RequestPerUrl:     []RequestsPerUrl{{Url: "any.com", Count: 4}, {Url: "abc.com", Count: 2}},
				HasData:           true,
			}},
		{"Interval 2 - no data",
			now.Add(time.Duration(10) * time.Second),
			[]ParsedPacket{},
			Stats{
				At:                "2020-01-01 00:00:10",
				AvgResponseTimeMs: 0,
				RequestPerUrl:     []RequestsPerUrl{},
				HasData:           false,
			}},
		{"Interval 3 - no data",
			now.Add(time.Duration(15) * time.Second),
			[]ParsedPacket{},
			Stats{
				At:                "2020-01-01 00:00:15",
				AvgResponseTimeMs: 0,
				RequestPerUrl:     []RequestsPerUrl{},
				HasData:           false,
			}},
		{"Interval 4",
			now.Add(time.Duration(20) * time.Second),
			[]ParsedPacket{
				{URL: "any.com", Uid: 100, At: now.Add(2 * time.Millisecond)},
				{URL: "any.com", Uid: 100, At: now.Add(4 * time.Millisecond)}, // 2 Millisecond resp time
			},
			Stats{
				At:                "2020-01-01 00:00:20",
				AvgResponseTimeMs: 2,
				RequestPerUrl:     []RequestsPerUrl{{Url: "any.com", Count: 2}},
				HasData:           true,
			}},
	}

	c := NewCalculator()
	for _, test := range tests {
		t.Run(test.title, func(t *testing.T) {
			for _, pack := range test.packets {
				c.ExtractPacketValues(pack)
			}

			actual := c.Stats(test.at).(Stats)
			assert.Equal(t, test.expected.At, actual.At)
			assert.Equal(t, test.expected.HasData, actual.HasData)
			assert.Equal(t, test.expected.AvgResponseTimeMs, actual.AvgResponseTimeMs)
			assert.Equal(t, len(test.expected.RequestPerUrl), len(actual.RequestPerUrl))

			for _, e := range test.expected.RequestPerUrl {
				for _, a := range actual.RequestPerUrl {
					if e.Url == a.Url {
						assert.Equal(t, e.Count, a.Count)
					}
				}
			}

			c.Cleanup()
		})
	}
}
