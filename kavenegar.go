package kavenegar

//Kavenegar ...
type Kavenegar struct {
	Message *MessageService
	Account *AccountService
	Verify  *VerifyService
	Call    *CallService
}

//New ...
func New(apikey string) *Kavenegar {
	client := NewClient(apikey)
	return NewWithClient(client)
}

//NewWithClient ...
func NewWithClient(client *Client) *Kavenegar {
	k := &Kavenegar{}
	k.Account = NewAccountService(client)
	k.Message = NewMessageService(client)
	k.Verify = NewVerifyService(client)
	k.Call = NewCallService(client)
	return k
}
