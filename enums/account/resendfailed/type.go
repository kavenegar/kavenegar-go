package AccountResendFailed

type Type string

const (
	Enabled  Type = "enabled"
	Disabled Type = "disabled"
)

var TypeMapName = map[Type]string{
	Enabled:  "enabled",
	Disabled: "disabled",
}

var TypeMapId = map[Type]string{
	Enabled:  "enabled",
	Disabled: "disabled",
}

func (t Type) String() string {
	return TypeMapId[t]
}
