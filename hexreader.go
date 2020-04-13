package hexreader

import (
	"errors"
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

func Decode(hs string) (*DataSeg, error) {
	//Check for an empty string
	if hs == "" {
		return nil, errors.New("[hexreader] Decode::Given an empty string")
	}
	//Remove the : character 
	//Convert the Hex Ascii to binary
	//Check the checksum

}