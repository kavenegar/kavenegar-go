package kavenegar

import (
	"time"
)

//Message ...
type Message struct {
	MessageID  int               `json:"messageid"`
	Message    string            `json:"message"`
	Status     MessageStatusType `json:"status"`
	StatusText string            `json:"statustext"`
	Sender     string            `json:"sender"`
	Receptor   string            `json:"receptor"`
	Date       int               `json:"date"`
	Cost       int               `json:"cost"`
}

//MessageSendParam ...
type MessageSendParam struct {
	Date    time.Time
	Type    []MessageSendType
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

type MessageSendType int

const (
	Type_MessageSend_Flash MessageSendType = iota
	Type_MessageSend_MobileMemory
	Type_MessageSend_SimMemory
	Type_MessageSend_AppMemory
)

var MessageSendMap = map[MessageSendType]string{
	Type_MessageSend_Flash:        "0",
	Type_MessageSend_MobileMemory: "1",
	Type_MessageSend_SimMemory:    "2",
	Type_MessageSend_AppMemory:    "3",
}

func (t MessageSendType) String() string {
	return MessageSendMap[t]
}

type MessageStatusType int

const (
	Type_MessageStatus_Queued       MessageStatusType = 1
	Type_MessageStatus_Schulded     MessageStatusType = 2
	Type_MessageStatus_SentToCenter MessageStatusType = 4
	Type_MessageStatus_Sent         MessageStatusType = 5
	Type_MessageStatus_Failed       MessageStatusType = 6
	Type_MessageStatus_Delivered    MessageStatusType = 10
	Type_MessageStatus_Undelivered  MessageStatusType = 11
	Type_MessageStatus_Canceled     MessageStatusType = 13
	Type_MessageStatus_Filtered     MessageStatusType = 14
	Type_MessageStatus_Received     MessageStatusType = 50
	Type_MessageStatus_Incorrect    MessageStatusType = 100
)

var MessageStatusMap = map[MessageStatusType]string{
	Type_MessageStatus_Queued:       "Queued",
	Type_MessageStatus_Schulded:     "Schulded",
	Type_MessageStatus_SentToCenter: "SentToCenter",
	Type_MessageStatus_Sent:         "Sent",
	Type_MessageStatus_Failed:       "Failed",
	Type_MessageStatus_Delivered:    "Delivered",
	Type_MessageStatus_Undelivered:  "Undelivered",
	Type_MessageStatus_Canceled:     "Canceled",
	Type_MessageStatus_Filtered:     "Filtered",
	Type_MessageStatus_Received:     "Received",
	Type_MessageStatus_Incorrect:    "Incorrect",
}

func (t MessageStatusType) String() string {
	return MessageStatusMap[t]
}
