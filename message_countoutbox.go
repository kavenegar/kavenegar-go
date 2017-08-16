package kavenegar

import (
	"net/url"
	"time"
)

//MessageCountOutbox ...
type MessageCountOutbox struct {
	*MessageCountInbox
	Sumpart int `json:"sumpart"`
	Cost    int `json:"cost"`
}

//MessageCountOutboxResult ...
type MessageCountOutboxResult struct {
	*Return `json:"return"`
	Entries []MessageCountOutbox `json:"entries"`
}

//CountOutbox ...
func (message *MessageService) CountOutbox(startdate time.Time, endate time.Time, status MessageStatusType) (MessageCountOutbox, error) {
	v := url.Values{}
	v.Set("startdate", ToUnix(startdate))
	if !endate.IsZero() {
		v.Set("endate", ToUnix(startdate))
	}
	v.Set("status", status.String())
	return message.CreateCountOutbox(v)
}

//CreateCountOutbox ...
func (message *MessageService) CreateCountOutbox(v url.Values) (MessageCountOutbox, error) {
	u := message.client.EndPoint("sms", "countoutbox")
	m := new(MessageCountOutboxResult)
	err := message.client.Execute(u.String(), v, m)
	if	m.Entries==nil{
		return MessageCountOutbox{},err
	}
	return m.Entries[0], err
}
