package kavenegar

import (
	"net/url"

	"github.com/kavenegar/kavenegar-go/enums/account/apilog"
	"github.com/kavenegar/kavenegar-go/enums/account/dailyreport"
	"github.com/kavenegar/kavenegar-go/enums/account/debugmode"
	"github.com/kavenegar/kavenegar-go/enums/account/resendfailed"
)

//AccountConfig ...
type AccountConfig struct {
	Apilogs        AccountAPILog.Type       `json:"apilogs"`
	Dailyreport    AccountDailyReport.Type  `json:"dailyreport"`
	Debugmode      AccountDebugMode.Type    `json:"debugmode"`
	Defaultsender  string                   `json:"defaultsender"`
	Mincreditalarm string                   `json:"mincreditalarm"`
	Resendfailed   AccountResendFailed.Type `json:"resendfailed"`
}

//AccountConfigParam ...
type AccountConfigParam struct {
	Apilogs        AccountAPILog.Type
	Dailyreport    AccountDailyReport.Type
	Debugmode      AccountDebugMode.Type
	Defaultsender  string
	Mincreditalarm string
	Resendfailed   AccountResendFailed.Type
}

//AccountConfigResult ...
type AccountConfigResult struct {
	*Return `json:"return"`
	Entries AccountConfig `json:"entries"`
}

//Config ...
func (c *AccountService) Config(param *AccountConfigParam) (AccountConfig, error) {
	v := structToURLValues(param)
	return c.CreateConfig(v)
}

//CreateConfig ..
func (c *AccountService) CreateConfig(v url.Values) (AccountConfig, error) {
	u := c.client.EndPoint("account", "config")
	m := new(AccountConfigResult)
	err := c.client.Execute(u.String(), v, m)
	return m.Entries, err
}
