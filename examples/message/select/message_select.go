package main

import (
	"fmt"

	"github.com/kavenegar/kavenegar-go"
)

func main() {
	api := kavenegar.NewKavenegar(" your apikey ")
	messageid := []int64{8792344,8792345}
	if res, err := api.Message.Select(messageid); err != nil {
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
