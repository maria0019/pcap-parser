package httppack

import (
	"time"
)

type ParsedPacket struct {
	URL string
	At  time.Time
	Uid int
}

func (p ParsedPacket) Url() string {
	return p.URL
}

func (p ParsedPacket) Timestamp() time.Time {
	return p.At
}

func (p ParsedPacket) ReqRespUid() int {
	return p.Uid
}

func (p ParsedPacket) IsValid() bool {
	return p.Uid != 0
}
