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
	c := NewClient(apikey)
	k := &Kavenegar{}
	k.Account = NewAccountService(c)
	k.Message = NewMessageService(c)
	k.Verify = NewVerifyService(c)
	k.Call = NewCallService(c)
	return k
}
