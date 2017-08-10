package kavenegar

import (
	"net/url"
	"time"
)

//SelectOutbox ...
func (message *MessageService) SelectOutbox(startdate time.Time, endate time.Time, sender string) ([]Message, error) {

	v := url.Values{}

	if !startdate.IsZero() {
		v.Set("startdate", ToUnix(startdate))
	}
	if !endate.IsZero() {
		v.Set("endate", ToUnix(endate))
	}
	if v.Get("sender") != "" {
		v.Set("sender", v.Get("sender"))
	}

	return message.CreateSelectOutbox(v)
}

//CreateSelectOutbox ...
func (message *MessageService) CreateSelectOutbox(v url.Values) ([]Message, error) {
	u := message.client.EndPoint("sms", "selectoutbox")
	m := new(MessageResult)
	err := message.client.Execute(u.String(), v, m)
	return m.Entries, err
}
