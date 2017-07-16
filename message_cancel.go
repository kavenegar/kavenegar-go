package kavenegar

import (
	"net/url"
)

//MessageCancel ...
type MessageCancel struct {
	Messageid  int    `json:"messageid"`
	Status     int    `json:"status"`
	Statustext string `json:"statustext"`
}

//MessageCancelResult ...
type MessageCancelResult struct {
	*Return `json:"return"`
	Entries []MessageCancel `json:"entries"`
}

//Cancel ...
func (message *MessageService) Cancel(messageid []int64) ([]MessageCancel, error) {
	u := message.client.EndPoint("sms", "cancel")
	m := new(MessageCancelResult)
	v := url.Values{}
	v.Set("messageid",ArrayEncodeToString(messageid))
	err := message.client.Execute(u.String(), v, m)
	return m.Entries, err
}
