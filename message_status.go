package kavenegar

import (
	"net/url"
)

//MessageStatus ...
type MessageStatus struct {
	MessageId  int    `json:"messageid"`
	Status     int    `json:"status"`
	StatusText string `json:"statustext"`
}

//MessageStatusResult ...
type MessageStatusResult struct {
	*Return `json:"return"`
	Entries []MessageStatus `json:"entries"`
}

//Status ...
func (message *MessageService) Status(messageid []string) ([]MessageStatus, error) {
	u := message.client.EndPoint("sms", "status")
	m := new(MessageStatusResult)
	v := url.Values{
		"messageid": messageid,
	}
	err := message.client.Execute(u.String(), v, m)
	return m.Entries, err
}
