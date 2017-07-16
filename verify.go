package kavenegar

import (
	"github.com/kavenegar/kavenegar-go/enums/verify/lookup"
)

//VerifyService ...
type VerifyService struct {
	client *Client
}

//VerifyLookupResult ...
type VerifyLookupResult struct {
	*Return `json:"return"`
	Entries Message `json:"entries"`
}

//VerifyLookupParam ...
type VerifyLookupParam struct {
	Token2 string
	Token3 string
	Type   VerifyLookup.Type
}

//NewVerifyService ...
func NewVerifyService(client *Client) *VerifyService {
	m := &VerifyService{client: client}
	return m
}
