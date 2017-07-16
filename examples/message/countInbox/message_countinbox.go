package main

import (
	"fmt"
	"net/url"
	"time"

	"github.com/kavenegar/kavenegar-go"
)

func main() {
	api := kavenegar.NewKavenegar(" your apikey ")

	//Message.CountInbox
	linenumber := ""
	startdate := time.Now().Add(time.Duration(24) * time.Hour)
	enddate := time.Now()
	isread := false
	if res, err := api.Message.CountInbox(linenumber, startdate, enddate, isread); err != nil {
		switch err := err.(type) {
		case *kavenegar.APIError:
			fmt.Println(err.Error())
		case *kavenegar.HTTPError:
			fmt.Println(err.Error())
		default:
			fmt.Println(err.Error())
		}
	} else {
		fmt.Println("Sumcount 	= ", res.Sumcount)
		fmt.Println("Startdate 	= ", res.Startdate)
		fmt.Println("Enddate 	= ", res.Enddate)
	}

	//Message.CreateCountInbox
	v := url.Values{}
	v.Set("linenumber", "")
	v.Set("startdate", kavenegar.TimeToUnix(time.Now().Add(time.Duration(24)*time.Hour)))
	v.Set("enddate", kavenegar.TimeToUnix(time.Now()))
	//v.Set("isread","1")
	if res, err := api.Message.CreateCountInbox(v); err != nil {
		switch err := err.(type) {
		case *kavenegar.APIError:
			fmt.Println(err.Error())
		case *kavenegar.HTTPError:
			fmt.Println(err.Error())
		default:
			fmt.Println(err.Error())
		}
	} else {
		fmt.Println("Sumcount 	= ", res.Sumcount)
		fmt.Println("Startdate 	= ", res.Startdate)
		fmt.Println("Enddate 	= ", res.Enddate)
	}
}
