package intelhex

import (
	"encoding/hex"
	"errors"
)

//Intel hex record type
type RecordType int

const (
	DataRecord	RecordType = iota // Data record type.
	EndOfFileRecord // End of file record.
	ExtendedSegmentAddressRecord // Extended segment address record.
	ExtendedLinearAddressRecord // Extended linear address record
	StartLinearAddressRecord // Start linear address record (MDK-ARM only).
	ErrorRecord // Used if error in a record.
)

// Contains info about one record.
type Record struct {
	RecordType RecordType // The type of record.
	Address uint32 // Address for bytes to be written to if data record.
	Data []byte // Data bytes if a data record.
}

func DecodeString(IntelHexString string) (*Record, error) {
	// Check for an empty string.
	if IntelHexString == "" {
		return &Record{
			RecordType: ErrorRecord,
			Address:    0,
			Data:       nil,
		}, errors.New("Given a blank string.")
	}

	// Check to see if intel hex string starts with a semi colon.
	if IntelHexString[0:1] != ":" {
		return &Record{
			RecordType: ErrorRecord,
			Address:    0,
			Data:       nil,
		}, errors.New("Intel Hex string does not start with :")
	}

	// Decode intel hex string into hex bytes.
	stringBytes, err := hex.DecodeString(IntelHexString[1:])
	if err != nil {
		return &Record{
			RecordType: ErrorRecord,
			Address:    0,
			Data:       nil,
		}, err
	}

	// Check to see if min number of bytes.
	if len(stringBytes) < 4 {
		return &Record{
			RecordType: ErrorRecord,
			Address:    0,
			Data:       nil,
		}, errors.New("Record does not have atleast four bytes.")
	}

	// Create a new record and make sure not nil.
	record := new(Record)
	if record == nil {
		return &Record{
			RecordType: ErrorRecord,
			Address:    0,
			Data:       nil,
		}, errors.New("Could not allocate space for a new record.")
	}

	switch RecordType(stringBytes[3]) {
	case DataRecord:
		record.RecordType = DataRecord

		recordLength := stringBytes[0]
		// Todo : Make sure correct number of bytes

		// Add the data bytes to the records slice.
		for i := 4; i < len(stringBytes); i++ {
			record.Data = append(record.Data, stringBytes[i])
		}

		return record, nil
	case EndOfFileRecord:
		record.RecordType = EndOfFileRecord
		return record, nil
	case ExtendedSegmentAddressRecord:
		record.RecordType = ExtendedSegmentAddressRecord
		return record, nil
	case ExtendedLinearAddressRecord:
		record.RecordType = ExtendedLinearAddressRecord
		return record, nil
	case StartLinearAddressRecord:
		record.RecordType = StartLinearAddressRecord
		return record, nil
	default:
		return &Record{
			RecordType: ErrorRecord,
			Address:    0,
			Data:       nil,
		}, errors.New("Not a valid record type.")
	}
}

