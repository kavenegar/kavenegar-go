package kavenegar

import (
	"net/url"
)

//MakeTTS ...
func (call *CallService) MakeTTS(receptor string, message string, params *CallParam) (Message, error) {
	v := structToURLValues(params)
	v.Set("receptor", receptor)
	v.Set("message", message)
	return call.CreateMakeTTS(v)
}

//CreateMakeTTS ...
func (call *CallService) CreateMakeTTS(v url.Values) (Message, error) {
	u := call.client.EndPoint("call", "maketts")
	res := new(MessageResult)
	err := call.client.Execute(u.String(), v, res)
	return res.Entries[0], err
}
