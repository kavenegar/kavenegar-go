package kavenegar

import (
	"time"

	"github.com/kavenegar/kavenegar-go/enums/message/send"
	"github.com/kavenegar/kavenegar-go/enums/message/status"
)

//Message ...
type Message struct {
	MessageID  int                    `json:"messageid"`
	Message    string                 `json:"message"`
	Status     MessageStatusType.Type `json:"status"`
	StatusText string                 `json:"statustext"`
	Sender     string                 `json:"sender"`
	Receptor   string                 `json:"receptor"`
	Date       int                    `json:"date"`
	Cost       int                    `json:"cost"`
}

//MessageParam ...
type MessageParam struct {
	Date    time.Time
	Type    []MessageSendType.Type
	LocalID []string
}

//MessageResult ...
type MessageResult struct {
	*Return
	Entries []Message `json:"entries"`
}

//MessageService ...
type MessageService struct {
	client *Client
}

//NewMessageService ...
func NewMessageService(client *Client) *MessageService {
	m := &MessageService{client: client}
	return m
}
