package kavenegar

import (
	"time"
)

//CallParam ...
type CallParam struct {
	Date    time.Time
	LocalID string
}

//CallService ...
type CallService struct {
	client *Client
}

//NewCallService ...
func NewCallService(client *Client) *CallService {
	m := &CallService{client: client}
	return m
}
