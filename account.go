package kavenegar

//AccountService ...
type AccountService struct {
	client *Client
}

//NewAccountService ...
func NewAccountService(client *Client) *AccountService {
	m := &AccountService{client: client}
	return m
}
