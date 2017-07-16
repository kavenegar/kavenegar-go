package kavenegar

import (
	"net/url"
	"strconv"
)

//MessageCountPostalCode ...
type MessageCountPostalCode struct {
	Section string `json:"section"`
	Value   int    `json:"value"`
}

//MessageCountPostalCodeResult ...
type MessageCountPostalCodeResult struct {
	*Return `json:"return"`
	Entries []MessageCountPostalCode `json:"entries"`
}

//CountPostalCode ...
func (message *MessageService) CountPostalCode(postalcode int64) ([]MessageCountPostalCode, error) {
	u := message.client.EndPoint("sms", "countpostalcode")
	m := new(MessageCountPostalCodeResult)
	v := url.Values{}
	v.Set("postalcode", strconv.FormatInt(postalcode, 10))
	err := message.client.Execute(u.String(), v, m)
	return m.Entries, err
}
