package kavenegar

import (
	"net/url"
)

//Send ...
func (m *MessageService) Send(sender string, receptor []string, message string, params *MessageSendParam) ([]Message, error) {
	v := url.Values{}
	v.Set("sender", sender)
	v.Set("receptor", ToString(receptor))
	v.Set("message", message)
	if params != nil {
		if !params.Date.IsZero() {
			v.Set("date", ToUnix(params.Date))
		}
		if params.Type != nil {
			v.Set("type", ToString(params.Type))
		}
		if params.LocalID != nil {
			v.Set("localid", ToString(params.LocalID))
		}
	}
	return m.CreateSend(v)
}

//CreateSend ...
func (m *MessageService) CreateSend(v url.Values) ([]Message, error) {
	u := m.client.EndPoint("sms", "send")
	res := new(MessageResult)
	err := m.client.Execute(u.String(), v, res)
	return res.Entries, err
}
