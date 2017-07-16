package main

import (
	"fmt"
	"net/url"
	"time"

	"github.com/kavenegar/kavenegar-go"
)

func main() {
	api := kavenegar.NewKavenegar(" your apikey ")

	startdate := time.Now().Add(time.Duration(10) * time.Minute)
	enddate := time.Now()
	sender := ""
	if res, err := api.Message.SelectOutbox(startdate, enddate, sender); err != nil {
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
	v.Set("startdate", kavenegar.TimeToUnix(t))
	v.Set("enddate", kavenegar.TimeToUnix(time.Now()))
	v.Set("sender", "")
	if res, err := api.Message.CreateSelectOutbox(v); err != nil {
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
