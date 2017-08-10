package kavenegar

//AccountInfo ...
type AccountInfo struct {
	Remaincredit int    `json:"remaincredit"`
	Expiredate   int    `json:"expiredate"`
	Type         string `json:"type"`
}

//AccountInfoResult ...
type AccountInfoResult struct {
	*Return `json:"return"`
	Entries AccountInfo `json:"entries"`
}

//Info ...
func (c *AccountService) Info() (AccountInfo, error) {
	u := c.client.EndPoint("account", "info")
	m := new(AccountInfoResult)
	err := c.client.Execute(u.String(), nil, m)
	return m.Entries, err
}
