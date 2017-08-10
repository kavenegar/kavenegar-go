package kavenegar

import (
	"net/url"
)

//MessageReceive ...
type MessageReceive struct {
	Messageid int    `json:"messageid"`
	Message   string `json:"message"`
	Sender    string `json:"sender"`
	Receptor  string `json:"receptor"`
	Date      int    `json:"date"`
}

//MessageReceiveResult ...
type MessageReceiveResult struct {
	*Return `json:"return"`
	Entries []MessageReceive `json:"entries"`
}

//Receive ...
func (message *MessageService) Receive(linenumber string, isread bool) ([]Message, error) {
	u := message.client.EndPoint("sms", "receive")
	m := new(MessageResult)
	v := url.Values{}
	v.Set("linenumber", linenumber)
	v.Set("isread", map[bool]string{true: "1", false: "0"}[isread != true])
	err := message.client.Execute(u.String(), v, m)
	return m.Entries, err
}
