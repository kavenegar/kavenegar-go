package kavenegar

import (
	"net/url"
	"strconv"
)

//MessageStatusLocal ...
type MessageStatusLocal struct {
	*MessageStatus
	LocalID string `json:"localid"`
}

//MessageStatusLocalResult ...
type MessageStatusLocalResult struct {
	*Return `json:"return"`
	Entries []MessageStatusLocal `json:"entries"`
}

//StatusLocal ...
func (message *MessageService) StatusLocal(localid int64) (MessageStatusLocal, error) {
	u := message.client.EndPoint("sms", "statuslocalmessageid")
	m := new(MessageStatusLocalResult)
	v := url.Values{}
	v.Set("localid", strconv.FormatInt(localid, 10))
	err := message.client.Execute(u.String(), v, m)
	if err!=nil{
		return MessageStatusLocal{}, err
	}
	return m.Entries[0], err
}
