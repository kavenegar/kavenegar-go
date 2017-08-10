package kavenegar

import (
	"net/url"
)

//AccountConfig ...
type AccountConfig struct {
	Apilogs        AccountAPILogType       `json:"apilogs"`
	Dailyreport    AccountDailyReportType  `json:"dailyreport"`
	Debugmode      AccountDebugModeType    `json:"debugmode"`
	Defaultsender  string                  `json:"defaultsender"`
	Mincreditalarm string                  `json:"mincreditalarm"`
	Resendfailed   AccountResendFailedType `json:"resendfailed"`
}

//AccountConfigParam ...
type AccountConfigParam struct {
	Apilogs        AccountAPILogType
	Dailyreport    AccountDailyReportType
	Debugmode      AccountDebugModeType
	Defaultsender  string
	Mincreditalarm string
	Resendfailed   AccountResendFailedType
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
	u := c.client.EndPoint("account", "info")
	m := new(AccountConfigResult)
	err := c.client.Execute(u.String(), v, m)
	return m.Entries, err
}
