package AccountAPILog

type Type string

const (
	Justfaults Type = "justfaults"
	Enabled    Type = "enabled"
	Disabled   Type = "disabled"
)

var TypeMapName = map[Type]string{
	Justfaults: "justfaults",
	Enabled:    "enabled",
	Disabled:   "disabled",
}

var TypeMapId = map[Type]string{
	Justfaults: "justfaults",
	Enabled:    "enabled",
	Disabled:   "disabled",
}

func (t Type) String() string {
	return TypeMapId[t]
}
