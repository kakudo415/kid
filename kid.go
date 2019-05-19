package kid

import (
	"fmt"
	"strconv"
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

// ToDec (Decimal num)
func (id ID) ToDec() string {
	return fmt.Sprintf("%020d", id)
}

// ToHex (Hexadecimal number)
func (id ID) ToHex(upper bool) string {
	if upper {
		return fmt.Sprintf("%016X", id)
	}
	return fmt.Sprintf("%016x", id)
}

// ToBin (Binary number)
func (id ID) ToBin() string {
	return fmt.Sprintf("%064b", id)
}

// ParseDec parses ID from decimal string
func ParseDec(src string) ID {
	id, err := strconv.ParseUint(src, 10, 64)
	if err != nil {
		return ID(0)
	}
	return ID(id)
}

// ParseHex parses ID from hexadecimal string
func ParseHex(src string) ID {
	id, err := strconv.ParseUint(src, 16, 64)
	if err != nil {
		return ID(0)
	}
	return ID(id)
}

// ParseBin parses ID from decimal string
func ParseBin(src string) ID {
	id, err := strconv.ParseUint(src, 2, 64)
	if err != nil {
		return ID(0)
	}
	return ID(id)
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
