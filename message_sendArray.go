package kavenegar

import (
	"net/url"
)

//SendArray ...
func (m *MessageService) SendArray(sender []string, receptor []string, message []string, params *MessageSendParam) ([]Message, error) {
	v := url.Values{}
	v.Set("sender", ToJson(sender))
	v.Set("receptor", ToJson(receptor))
	v.Set("message", ToJson(message))
	if params != nil {
		if !params.Date.IsZero() {
			v.Set("date", ToUnix(params.Date))
		}
		if params.Type != nil {
			v.Set("type",  ToJson(params.Type))
		}
		if params.LocalID != nil {
			v.Set("localid",  ToJson(params.LocalID))
		}
	}
	return m.CreateSendArray(v)
}

//CreateSendArray ...
func (m *MessageService) CreateSendArray(v url.Values) ([]Message, error) {
	u := m.client.EndPoint("sms", "sendarray")
	res := new(MessageResult)
	err := m.client.Execute(u.String(), v, res)
	return res.Entries, err
}
