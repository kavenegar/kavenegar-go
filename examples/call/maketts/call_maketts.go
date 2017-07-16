package main

import (
	"fmt"
	"net/url"

	"github.com/kavenegar/kavenegar-go"
)

func main() {
	api := kavenegar.NewKavenegar(" your apikey ")

	//Call.MakeTTS
	receptor := ""
	message := ""
	params := &kavenegar.CallParam{}
	if res, err := api.Call.MakeTTS(receptor, message, params); err != nil {
		switch err := err.(type) {
		case *kavenegar.APIError:
			fmt.Println(err.Error())
		case *kavenegar.HTTPError:
			fmt.Println(err.Error())
		default:
			fmt.Println(err.Error())
		}
	} else {
		fmt.Println("MessageID 	= ", res.MessageID)
		fmt.Println("Status    	= ", res.Status)
		//...
	}

	//Call.CreateMakeTTS
	v := url.Values{}
	v.Set("receptor", "")
	v.Add("message", "")
	//v.Add("date", "")
	//v.Add("localid", "")
	if res, err := api.Call.CreateMakeTTS(v); err != nil {
		switch err := err.(type) {
		case *kavenegar.APIError:
			fmt.Println(err.Error())
		case *kavenegar.HTTPError:
			fmt.Println(err.Error())
		default:
			fmt.Println(err.Error())
		}
	} else {
		fmt.Println("MessageID 	= ", res.MessageID)
		fmt.Println("Status    	= ", res.Status)
		//...
	}
}
