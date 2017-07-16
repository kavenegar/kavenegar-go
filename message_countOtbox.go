package kavenegar

import (
	"net/url"
	"time"

	"github.com/kavenegar/kavenegar-go/enums/message/status"
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
func (message *MessageService) CountOutbox(startdate time.Time, endate time.Time, status MessageStatusType.Type) (MessageCountOutbox, error) {
	v := url.Values{}
	v.Set("startdate", TimeToUnix(startdate))
	if !endate.IsZero() {
		v.Set("endate", TimeToUnix(startdate))
	}
	v.Set("status", status.String())
	return message.CreateCountOutbox(v)
}

//CreateCountOutbox ...
func (message *MessageService) CreateCountOutbox(v url.Values) (MessageCountOutbox, error) {
	u := message.client.EndPoint("sms", "countoutbox")
	m := new(MessageCountOutboxResult)
	err := message.client.Execute(u.String(), v, m)
	if err!=nil{
		return MessageCountOutbox{}, err
	}
	return m.Entries[0], err
}
