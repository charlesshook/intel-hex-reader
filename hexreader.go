package hexreader

import (
	"errors"
	//"encoding/hex"
)

//Intel hex record type
type RecType int

/*
	DATA - Data record type
	EOF - End of file record type
	EST - Extended segment address record type
	SSA - Start segment address record type
	ELA - Extended linear address record type
	SLA - Start linear address
*/
const (
	DATA	RecType = iota
	EOF	
	EST
	SSA
	ELA
	SLA
)

type DataSeg struct {
	Adr uint32
	Data []byte
}

type Rec struct {
	rType RecType
	Adr uint16
	Data []byte
}

func Decode(hs string) (*Rec, error) {
	//Check for an empty string
	if hs == "" {
		return nil, errors.New("[hexreader] Decode::Given an empty string")
	}

	//Remove the : character 
	hs = hs[1:]

	//Convert the Hex Ascii to binary
	/*
	b, err := hex.DecodeString(hs)
	if err != nil {
		return nil, errors.New("[hexreader] Decode::Unable to decode hex string.")
	}
	*/

	//Check the checksum
	
	//Allocate memory for a new record
	rec := new(Rec)

	return rec, nil
}

func calcChecksum(b []byte) uint8 {
	var s int

	//Sum up all the data bytes
	for _, n := range b {
		s += int(n)
	}

	return byte(-s)
}