package kavenegar

// VerifyService ...
type VerifyService struct {
	client *Client
}

// VerifyLookupResult ...
type VerifyLookupResult struct {
	*Return `json:"return"`
	Entries Message `json:"entries"`
}

// VerifyLookupParam ...
type VerifyLookupParam struct {
	Token2 string
	Token3 string
	Tokens map[string]string
	Type   VerifyLookupType
}

// NewVerifyService ...
func NewVerifyService(client *Client) *VerifyService {
	m := &VerifyService{client: client}
	return m
}

type VerifyLookupType string

const (
	Type_VerifyLookup_Sms  VerifyLookupType = "sms"
	Type_VerifyLookup_Call VerifyLookupType = "call"
)
