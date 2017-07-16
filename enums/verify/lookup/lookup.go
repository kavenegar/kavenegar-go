package VerifyLookup

type Type int

const (
	Sms Type = iota
	Call
)

var TypeMapName = map[Type]string{
	Sms:  "Sms",
	Call: "Call",
}

var TypeMapId = map[Type]string{
	Sms:  "0",
	Call: "1",
}

func (t Type) String() string {
	return TypeMapId[t]
}
