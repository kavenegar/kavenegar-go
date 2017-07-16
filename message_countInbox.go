package kavenegar

import (
	"net/url"
	"time"
)

//MessageCountInbox ...
type MessageCountInbox struct {
	Startdate int `json:"startdate"`
	Enddate   int `json:"enddate"`
	Sumcount  int `json:"sumcount"`
}

//MessageCountInboxResult ...
type MessageCountInboxResult struct {
	*Return `json:"return"`
	Entries []MessageCountInbox `json:"entries"`
}

//CountInbox ...
func (message *MessageService) CountInbox(linenumber string, startdate time.Time, endate time.Time, isread bool) (MessageCountInbox, error) {

	v := url.Values{}
	v.Set("linenumber", linenumber)
	v.Set("startdate", TimeToUnix(startdate))
	if !endate.IsZero() {
		v.Set("endate", TimeToUnix(startdate))
	}
	v.Set("isread", map[bool]string{true: "1", false: "0"}[isread != true])
	return message.CreateCountInbox(v)
}

//CreateCountInbox ...
func (message *MessageService) CreateCountInbox(v url.Values) (MessageCountInbox, error) {
	u := message.client.EndPoint("sms", "countinbox")
	m := new(MessageCountInboxResult)
	err := message.client.Execute(u.String(), v, m)
	if err!=nil{
		return MessageCountInbox{}, err
	}
	return m.Entries[0], err
}
