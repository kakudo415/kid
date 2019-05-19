package kid

import (
	"strconv"
	"testing"
)

func Test_NewID(t *testing.T) {
	if id := New(255); id == 0 {
		t.Error("INPUT VALID WORKER ID, BUT ID == 0")
	} else {
		t.Log(id)
	}

	if id := New(256); id != 0 {
		t.Error("INPUT INVALID WORKER ID, BUT RETURNED ID")
	}

	if id := New(256); !id.IsError() {
		t.Error("INPUT INVALID WORKER ID, BUT ID ISN'T ERROR")
	}
}

func Test_SequenceNumber(t *testing.T) {
	sequenceNumber = 4090
	for i := 0; i < 10; i++ {
		if n := newSequenceNumber(); n <= 0xFFF {
			t.Log("SEQUENCE NUMBER VALID " + strconv.FormatUint(n, 10))
		} else {
			t.Error("SEQUENCE NUMBER OVERFLOW " + strconv.FormatUint(n, 10))
			break
		}
	}
}
