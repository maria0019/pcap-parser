package mock

import (
	"pparse/mock"
	"time"
)

type ParsedPacketValid struct{}

func (p ParsedPacketValid) Timestamp() time.Time {
	t, _ := time.Parse(time.DateTime, mock.TimeNow)

	return t
}

func (p ParsedPacketValid) IsValid() bool {
	return true
}

type ParsedPacketInvalid struct{}

func (p ParsedPacketInvalid) Timestamp() time.Time {
	t, _ := time.Parse(time.DateTime, mock.TimeNow)

	return t
}

func (p ParsedPacketInvalid) IsValid() bool {
	return false
}
