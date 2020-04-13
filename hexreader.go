package hexreader

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