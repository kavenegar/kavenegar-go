package main

import (
	"fmt"
	"net/url"

	"github.com/kavenegar/kavenegar-go"
)

func main() {
	api := kavenegar.NewKavenegar(" your apikey ")

	//Verify.VerifyLookup
	receptor := ""
	template := ""
	token := ""
	params := &kavenegar.VerifyLookupParam{}
	if res, err := api.Verify.Lookup(receptor, template, token, params); err != nil {
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

	//Verify.CreateLookup
	v := url.Values{}
	v.Set("receptor", "")
	v.Add("template", "")
	v.Add("token", "")
	//v.Add("token2", "")
	//v.Add("token3", "")
	if res, err := api.Verify.CreateLookup(v); err != nil {
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
