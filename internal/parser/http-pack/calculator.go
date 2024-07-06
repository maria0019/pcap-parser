package httppack

import (
	"encoding/json"
	"pparse/internal/calculator"
	"pparse/internal/packet"
	"strings"
	"time"
)

type Calculator struct {
	RequestsPerUrl map[string]int
	ResponseTimes  map[int][]time.Time
}

type Stats struct {
	At                string
	AvgResponseTimeMs int64
	RequestPerUrl     []RequestsPerUrl
	HasData           bool
}

type RequestsPerUrl struct {
	Url   string `json:"url"`
	Count int    `json:"count"`
}

func NewCalculator() calculator.CalculatorI {
	c := Calculator{
		RequestsPerUrl: map[string]int{},
		ResponseTimes:  map[int][]time.Time{},
	}

	return &c
}

func (c *Calculator) ExtractPacketValues(p packet.ParsedPacketI) {
	pack := p.(ParsedPacket)

	c.addResponseTime(pack.ReqRespUid(), pack.Timestamp())
	c.addRequestsPerUrlToPrint(pack.Url())

}

func (c *Calculator) Stats(at time.Time) interface{} {
	return Stats{
		At:                at.Format(time.DateTime),
		AvgResponseTimeMs: c.calculateAverageResponseTime(),
		RequestPerUrl:     c.requestsPerUrlToPrint(),
		HasData:           len(c.requestsPerUrlToPrint()) > 0,
	}
}

func (c *Calculator) StatsAsMap(at time.Time) map[string]interface{} {
	stats := c.Stats(at).(Stats)
	b, _ := json.Marshal(stats.RequestPerUrl)

	return map[string]interface{}{
		"at":                stats.At,
		"avgResponseTimeMs": stats.AvgResponseTimeMs,
		"requestPerUrl":     string(b),
		"hasData":           stats.HasData,
	}
}

func (c *Calculator) Cleanup() {
	c.RequestsPerUrl = map[string]int{}
	c.ResponseTimes = map[int][]time.Time{}
}

// CalculateAverageResponseTime Milliseconds
func (c *Calculator) calculateAverageResponseTime() int64 {
	var (
		sum   int64
		count int64
	)

	for _, startEndTimes := range c.ResponseTimes {
		if len(startEndTimes) < 2 {
			continue // may be response is not received yet
		}

		count++
		sum += startEndTimes[1].Sub(startEndTimes[0]).Milliseconds()
	}

	if count == 0 {
		return 0
	}

	return sum / count
}

func (c *Calculator) requestsPerUrlToPrint() []RequestsPerUrl {
	if len(c.RequestsPerUrl) == 0 {
		return []RequestsPerUrl{}
	}

	res := make([]RequestsPerUrl, 0)
	for url, val := range c.RequestsPerUrl {
		res = append(res, RequestsPerUrl{
			Url:   url,
			Count: val,
		})
	}
	return res
}

func (c *Calculator) addResponseTime(uid int, timestamp time.Time) {
	if _, ok := c.ResponseTimes[uid]; !ok {
		c.ResponseTimes[uid] = []time.Time{}
	}

	c.ResponseTimes[uid] = append(c.ResponseTimes[uid], timestamp)
}

func (c *Calculator) addRequestsPerUrlToPrint(url string) {
	if url == "" {
		return
	}

	if strings.HasSuffix(url, "/") {
		url = url[:len(url)-len("/")]
	}

	if _, ok := c.RequestsPerUrl[url]; !ok {
		c.RequestsPerUrl[url] = 0
	}

	c.RequestsPerUrl[url]++
}
