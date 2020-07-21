package hexreader

import (
	"testing"
)

func TestDecodeString(t *testing.T) {
	hexreader := New()

	dataRecordString := ":10246200464C5549442050524F46494C4500464C33"

	record, _ := hexreader.DecodeString(dataRecordString)

	if record.RecordType != DataRecord {
		t.Errorf("Wanted a DataRecord but got [%d]", record.RecordType)
	}
}
