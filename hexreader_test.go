package hexreader

import (
	"testing"
)

func TestDecode(t *testing.T) {
	s := ":04000005000000CD2A"

	_, err := Decode(s)
	if err != nil {
		t.Errorf("Received error [%s]", err)
	}
}