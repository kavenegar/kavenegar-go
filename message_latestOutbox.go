package kavenegar

import (
	"net/url"
	"strconv"
)

//LatestOutbox ...
func (message *MessageService) LatestOutbox(sender string, pagesize int) ([]Message, error) {
	v := url.Values{}
	v.Set("sender", sender)
	v.Set("pagesize", strconv.Itoa(pagesize))
	return message.CreateLatestOutbox(v)
}

//CreateLatestOutbox ...
func (message *MessageService) CreateLatestOutbox(v url.Values) ([]Message, error) {
	u := message.client.EndPoint("sms", "latestoutbox")
	vc := url.Values{}
	if v.Get("sender") != "" {
		vc.Set("sender", v.Get("sender"))
	}
	if v.Get("pagesize") != "" {
		vc.Set("pagesize", v.Get("pagesize"))
	}
	m := new(MessageResult)
	err := message.client.Execute(u.String(), v, m)
	return m.Entries, err
}
