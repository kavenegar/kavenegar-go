package main

import (
	"fmt"
	"net/url"

	"github.com/kavenegar/kavenegar-go"
)

func main() {
	api := kavenegar.NewKavenegar(" your apikey ")

	//Message.SendArray

	receptor := []string{"", ""} //Recipient numbers
	message := []string{"", ""}  //Text messages
	sender := []string{"", ""}   //Sender Line Numbers(optional)

	if res, err := api.Message.SendArray(sender, receptor, message, nil); err != nil {
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

	//Message.CreateSendArray
	v := url.Values{}
	v.Set("sender", kavenegar.ArrayEncodeToString([]string{"",""}))
	v.Set("message", kavenegar.ArrayEncodeToString([]string{"Hello Go!","Hello Go!"}))
	v.Set("receptor", kavenegar.ArrayEncodeToString([]string{"", ""}))
	//v.Set("type", kavenegar.ArrayEncodeToString([]string{sendType.AppMemory.String(), sendType.AppMemory.String()}))
	//v.Set("localid", kavenegar.ArrayEncodeToString([]string{"1000","1001"}))
	//t := time.Now().Add(time.Duration(10) * time.Minute)
	//v.Set("date", kavenegar.TimeToUnix(t))
	if res, err := api.Message.CreateSendArray(v); err != nil {
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
