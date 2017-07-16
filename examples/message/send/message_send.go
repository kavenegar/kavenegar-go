package main

import (
	"fmt"
	"net/url"
	"time"

	"github.com/kavenegar/kavenegar-go"
	"github.com/kavenegar/kavenegar-go/enums/message/send"
)

func main() {
	api := kavenegar.NewKavenegar(" your apikey ")

	//Message.Send
	sender := ""                 //Sender Line Number(optional)
	receptor := []string{"", ""} //Recipient numbers
	message := "Hello Go!"       //Text message
	params := &kavenegar.MessageParam{
		Date:    time.Now().Add(time.Duration(10) * time.Minute),
		LocalID: []string{"1000", "1001"},
		Type:    []MessageSendType.Type{MessageSendType.AppMemory, MessageSendType.AppMemory},
	}
	if res, err := api.Message.Send(sender, receptor, message, params); err != nil {
		switch err := err.(type) {
		case *kavenegar.APIError:
			fmt.Println(err.Error())
		case *kavenegar.HTTPError:
			fmt.Println(err.Error())
		default:
			fmt.Println(err.Error())
		}
	} else {
		for _, r := range res {
			fmt.Println("MessageID 	= ", r.MessageID)
			fmt.Println("Status    	= ", r.Status)
			//...
		}
	}

	//Message.CreateSend
	v := url.Values{}
	//v.Set("sender", "")
	v.Set("message", "Hello Go!")
	v.Set("receptor", kavenegar.ArrayEncodeToString([]string{"", ""}))
	//v.Set("type", kavenegar.ArrayEncodeToString([]string{sendType.AppMemory.String(), sendType.AppMemory.String()}))
	//v.Set("localid", kavenegar.ArrayEncodeToString([]string{"1000","1001"}))
	//t := time.Now().Add(time.Duration(10) * time.Minute)
	//v.Set("date", kavenegar.TimeToUnix(t))
	if res, err := api.Message.CreateSend(v); err != nil {
		switch err := err.(type) {
		case *kavenegar.APIError:
			fmt.Println(err.Error())
		case *kavenegar.HTTPError:
			fmt.Println(err.Error())
		default:
			fmt.Println(err.Error())
		}
	} else {
		for _, r := range res {
			fmt.Println("MessageID 	= ", r.MessageID)
			fmt.Println("Status    	= ", r.Status)
			//...
		}
	}
}
