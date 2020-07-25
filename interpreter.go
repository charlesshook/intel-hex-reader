package intelhex

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

type Interpreter struct {

}

func New() *Interpreter {
	return new(Interpreter)
}

func calculateChecksum(bytes []byte) (checksum byte) {
	for i := 0; i < len(bytes); i++ {
		checksum += bytes[i]
	}

	checksum += 0x01

	// Flip the bits
	checksum = 1 * -checksum

	return checksum
}