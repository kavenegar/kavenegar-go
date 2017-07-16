package kavenegar

import (
	"net/url"
)

//MessageStatus ...
type MessageStatus struct {
	Messageid  int    `json:"messageid"`
	Status     int    `json:"status"`
	Statustext string `json:"statustext"`
}

//MessageStatusResult ...
type MessageStatusResult struct {
	*Return `json:"return"`
	Entries []MessageStatus `json:"entries"`
}

//Status ...
func (message *MessageService) Status(messageid []int64) ([]MessageStatus, error) {
	u := message.client.EndPoint("sms", "status")
	m := new(MessageStatusResult)
	v := url.Values{}
	v.Set("messageid",ArrayEncodeToString(messageid))
	err := message.client.Execute(u.String(), v, m)
	return m.Entries, err
}
