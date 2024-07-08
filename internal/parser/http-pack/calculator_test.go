package httppack

import (
	"github.com/stretchr/testify/assert"
	"pparse/mock"
	"testing"
	"time"
)

func Test_Calculator(t *testing.T) {
	now, _ := time.Parse(time.DateTime, mock.TimeNow)
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
				At:                mock.TimeNowAdd5Sec,
				AvgResponseTimeMs: 6,
				RequestPerUrl:     []RequestsPerUrl{{Url: "any.com", Count: 4}, {Url: "abc.com", Count: 2}},
				HasData:           true,
			}},
		{"Interval 2 - no data",
			now.Add(time.Duration(10) * time.Second),
			[]ParsedPacket{},
			Stats{
				At:                mock.TimeNowAdd10Sec,
				AvgResponseTimeMs: 0,
				RequestPerUrl:     []RequestsPerUrl{},
				HasData:           false,
			}},
		{"Interval 3 - no data",
			now.Add(time.Duration(15) * time.Second),
			[]ParsedPacket{},
			Stats{
				At:                mock.TimeNowAdd15Sec,
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
				At:                mock.TimeNowAdd20Sec,
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

			actualStats := c.Stats(test.at).(Stats)
			assert.Equal(t, test.expected.At, actualStats.At)
			assert.Equal(t, test.expected.HasData, actualStats.HasData)
			assert.Equal(t, test.expected.AvgResponseTimeMs, actualStats.AvgResponseTimeMs)
			assert.Equal(t, len(test.expected.RequestPerUrl), len(actualStats.RequestPerUrl))

			for _, e := range test.expected.RequestPerUrl {
				for _, a := range actualStats.RequestPerUrl {
					if e.Url == a.Url {
						assert.Equal(t, e.Count, a.Count)
					}
				}
			}

			actualStatsAsMap := c.StatsAsMap(test.at)
			assert.Equal(t, test.expected.At, actualStatsAsMap["at"])
			assert.Equal(t, test.expected.HasData, actualStatsAsMap["hasData"])
			assert.Equal(t, test.expected.AvgResponseTimeMs, actualStatsAsMap["avgResponseTimeMs"])
			assert.NotEmpty(t, actualStatsAsMap["requestPerUrl"])

			c.Cleanup()
		})
	}
}
