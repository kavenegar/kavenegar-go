package MessageStatusType

type Type int

const (
	Queued       Type = 1
	Schulded     Type = 2
	SentToCenter Type = 4
	Delivered    Type = 10
	Undelivered  Type = 11
	Canceled     Type = 13
	Filtered     Type = 14
	Received     Type = 50
	Incorrect    Type = 100
)

var TypeMap = map[Type]string{
	Queued:       "Queued",
	Schulded:     "Schulded",
	SentToCenter: "SentToCenter",
	Delivered:    "Delivered",
	Undelivered:  "Undelivered",
	Canceled:     "Canceled",
	Filtered:     "Filtered",
	Received:     "Received",
	Incorrect:    "Incorrect",
}

func (t Type) String() string {
	return TypeMap[t]
}
