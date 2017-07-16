package kavenegar

import (
	"net/url"
)

//Select ...
func (message *MessageService) Select(messageid []int64) ([]Message, error) {
	u := message.client.EndPoint("sms", "select")
	m := new(MessageResult)
	v := url.Values{}
	v.Set("messageid",ArrayEncodeToString(messageid))
	err := message.client.Execute(u.String(), v, m)
	return m.Entries, err
}
