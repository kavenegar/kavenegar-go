package main

import (
	"fmt"
	"net/url"
	"time"

	"github.com/kavenegar/kavenegar-go"
)

func main() {
	api := kavenegar.NewKavenegar(" your apikey ")

	var postalcode int64 = 141
	var sender = ""
	var message = ""
	var mcistartindex = 1
	var mcicount = 1
	var mtnstartindex = 1
	var mtncount = 1
	var date = time.Now().Add(time.Duration(10) * time.Minute)

	if res, err := api.Message.SendPostalCode(postalcode, sender, message, mcistartindex, mcicount, mtnstartindex, mtncount, date); err != nil {
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

	v := url.Values{}
	t := time.Now().Add(time.Duration(10) * time.Minute)
	v.Set("postalcode", "")
	v.Set("sender", "")
	v.Set("message", "")
	v.Set("mcistartindex", "")
	v.Set("mcicount", "")
	v.Set("mtnstartindex ", "")
	v.Set("mtncount", "")
	v.Set("date ", kavenegar.TimeToUnix(t))
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
