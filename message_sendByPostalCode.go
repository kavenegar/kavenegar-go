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
	v.Set("date", TimeToUnix(date))
	return m.CreateSendPostalCode(v)
}

//CreateSendPostalCode ...
func (m *MessageService) CreateSendPostalCode(v url.Values) ([]Message, error) {
	u := m.client.EndPoint("sms", "sendbypostalcode")
	vc := url.Values{}
	if v.Get("postalcode") != "" {
		vc.Set("postalcode", v.Get("postalcode"))
	}
	if v.Get("sender") != "" {
		vc.Set("sender", v.Get("sender"))
	}
	if v.Get("message") != "" {
		vc.Set("message", v.Get("message"))
	}
	if v.Get("mcistartind") != "" {
		vc.Set("mcistartind", v.Get("mcistartind"))
	}
	if v.Get("mcicount") != "" {
		vc.Set("mcicount", v.Get("mcicount"))
	}
	if v.Get("mtnstartind") != "" {
		vc.Set("mtnstartind", v.Get("mtnstartind"))
	}
	if v.Get("mtncount") != "" {
		vc.Set("mtncount", v.Get("mtncount"))
	}
	if v.Get("date") != "" {
		vc.Set("date", v.Get("date"))
	}
	res := new(MessageResult)
	err := m.client.Execute(u.String(), v, res)
	return res.Entries, err
}
