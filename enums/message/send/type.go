package MessageSendType

type Type int

const (
	Flash Type = iota
	MobileMemory
	SimMemory
	AppMemory
)

var TypeMapName = map[Type]string{
	Flash:        "Flash",
	MobileMemory: "MobileMemory",
	SimMemory:    "SimMemory",
	AppMemory:    "AppMemory",
}

var TypeMapId = map[Type]string{
	Flash:        "0",
	MobileMemory: "1",
	SimMemory:    "2",
	AppMemory:    "3",
}

func (t Type) String() string {
	return TypeMapId[t]
}
