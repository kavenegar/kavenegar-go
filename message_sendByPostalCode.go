package kavenegar

import (
	"net/url"
	"strconv"
	"time"
)

//SendPostalCode ...
func (m *MessageService) SendPostalCode(postalcode int64, sender string, message string, mcistartindex int, mcicount int, mtnstartindex int, mtncount int, date time.Time) ([]Message, error) {
	v := url.Values{}
	v.Set("postalcode", strconv.FormatInt(postalcode, 10))
	v.Set("sender", sender)
	v.Set("message", message)
	v.Set("mcistartind", strconv.Itoa(mcistartindex))
	v.Set("mcicount", strconv.Itoa(mcicount))
	v.Set("mtnstartind", strconv.Itoa(mtnstartindex))
	v.Set("mtncount", strconv.Itoa(mtncount))
	v.Set("date", ToUnix(date))
	return m.CreateSendPostalCode(v)
}

//CreateSendPostalCode ...
func (m *MessageService) CreateSendPostalCode(v url.Values) ([]Message, error) {
	u := m.client.EndPoint("sms", "sendbypostalcode")	
	res := new(MessageResult)
	err := m.client.Execute(u.String(), v, res)
	return res.Entries, err
}
