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
	sequenceNumber = 4093
	for i := 0; i < 5; i++ {
		if n := newSequenceNumber(); n <= 0xFFF {
			t.Log("SEQUENCE NUMBER VALID " + strconv.FormatUint(n, 10))
		} else {
			t.Error("SEQUENCE NUMBER OVERFLOW " + strconv.FormatUint(n, 10))
			break
		}
	}
}

func Test_ToDec(t *testing.T) {
	d := New(0).ToDec()
	t.Log(d)
	if len(d) != 20 {
		t.Error("NUMBER OF DIGITS ISN'T 20 (Dec num)")
	}
}

func Test_ToHex(t *testing.T) {
	us := New(0).ToHex(true)
	t.Log(us)
	ls := New(0).ToHex(false)
	t.Log(ls)
	if len(us) != 16 || len(ls) != 16 {
		t.Error("NUMBER OF DIGITS ISN'T 16 (Hex num)")
	}
}

func Test_ToBin(t *testing.T) {
	b := New(0).ToBin()
	t.Log(b)
	if len(b) != 64 {
		t.Error("NUMBER OF DIGITS ISN'T 64 (Bin num)")
	}
}
