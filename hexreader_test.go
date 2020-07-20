package hexreader

import (
	"testing"
)

func TestDecodeString(t *testing.T) {
	hexreader := New()

	dataRecordString := ":10246200464C5549442050524F46494C4500464C33"

	record, err := hexreader.DecodeString(dataRecordString)

	if record. != got {
		t.Error("Wrong")
	}
}
