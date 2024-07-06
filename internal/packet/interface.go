package packet

import "time"

type ParsedPacketI interface {
	Timestamp() time.Time
	IsValid() bool
}
