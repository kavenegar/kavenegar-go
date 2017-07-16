package main

import (
	"fmt"

	"github.com/kavenegar/kavenegar-go"
)

func main() {
	api := kavenegar.NewKavenegar(" your apikey ")
	var postalcode int64 = 11
	if res, err := api.Message.CountPostalCode(postalcode); err != nil {
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
			fmt.Println("Section 	= ", r.Section)
			fmt.Println("Value    	= ", r.Value)
			//...
		}
	}
}
