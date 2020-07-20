package hexreader

import (
	//"encoding/hex"
	"errors"
)

//Intel hex record type
type RecordType int

/*
	DATA - Data record type
	EOF - End of file record type
	EST - Extended segment address record type
	SSA - Start segment address record type
	ELA - Extended linear address record type
	SLA - Start linear address
*/
const (
	DATA	RecordType = iota
	EOF	
	EST
	SSA
	ELA
	SLA
	ERROR
)

type IntelHexReader interface {
	DecodeString(IntelHexString string) (*Record, error)
	// DecodeBytes(IntelHexBytes []byte) (*Record, RecordType, error)
	// DecodeFile(IntelHexFile string) (*[]Record, error)
}

type IntelHex struct {

}

type Record struct {
	RecordType RecordType
	Address uint32
	Data []byte
}

func New() IntelHexReader {
	return new(IntelHex)
}

func (IH *IntelHex) DecodeString(IntelHexString string) (*Record, error) {
	// Check to see if intel hex string starts with a semi colon
	if IntelHexString[1:] != ":" {
		return nil, errors.New("Intel Hex string does not start with :")
	}

	// Check to see if intel hex string has enough bytes

	temp := new(Record)
	return temp, nil
}