package kid

import (
	"time"
)

// ID [64b]
type ID uint64

// New KID [64b]
// Worker ID is 8b (<= 255)
func New(workerID uint64) ID {
	if workerID > 0xFF {
		return 0
	}
	return ID(newTimestamp()<<20 + newSequenceNumber()<<8 + workerID)
}

// IsError (ID == 0 is Error)
func (id ID) IsError() bool {
	return id == 0
}

// Timestamp [44b]
var epoch = time.Date(2000, time.January, 1, 0, 0, 0, 0, time.UTC) // 2000-1-1T00:00:00Z

func newTimestamp() uint64 {
	elapsed := time.Now().UTC().Sub(epoch)
	timestamp := uint64(elapsed.Nanoseconds()) / uint64(time.Millisecond)
	return timestamp
}

// Sequence Number [12b]
var sequenceNumber uint64

const mask uint64 = 0xFFF

func newSequenceNumber() uint64 {
	sequenceNumber++
	if sequenceNumber > 0xFFF {
		sequenceNumber = 0
	}
	return sequenceNumber & mask
}

func init() {
	if !time.Now().UTC().After(epoch) {
		panic("KID ONLY WORKS AFTER 2000-1-1T00:00:00")
	}
}
