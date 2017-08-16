package kavenegar

import (
	"net/url"
)

//Lookup ...
func (c *VerifyService) Lookup(receptor string, template string, token string, param *VerifyLookupParam) (Message, error) {
	v := structToURLValues(param)
	v.Set("receptor", receptor)
	v.Set("template", template)
	v.Set("token", token)
	return c.CreateLookup(v)
}

//CreateLookup ..
func (c *VerifyService) CreateLookup(v url.Values) (Message, error) {
	u := c.client.EndPoint("verify", "lookup")
	m := new(MessageResult)
	err := c.client.Execute(u.String(), v, m)
	if	m.Entries==nil{
		return Message{},err
	}
	return m.Entries[0], err
}
