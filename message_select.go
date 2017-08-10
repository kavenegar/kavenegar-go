package kavenegar

import (
	"net/url"
)

//Select ...
func (message *MessageService) Select(messageid []string) ([]Message, error) {
	u := message.client.EndPoint("sms", "select")
	m := new(MessageResult)
	v := url.Values{
		"messageid": messageid,
	}
	err := message.client.Execute(u.String(), v, m)
	return m.Entries, err
}
