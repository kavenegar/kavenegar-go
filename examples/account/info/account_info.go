package main

import (
	"fmt"

	"github.com/kavenegar/kavenegar-go"
)

func main() {
	api := kavenegar.NewKavenegar(" your apikey ")
	if res, err := api.Account.Info(); err != nil {
		switch err := err.(type) {
		case *kavenegar.APIError:
			fmt.Println(err.Error())
		case *kavenegar.HTTPError:
			fmt.Println(err.Error())
		default:
			fmt.Println(err.Error())
		}
	} else {
		fmt.Println("Expiredate 	= ", res.Expiredate)
		fmt.Println("Remaincredit 	= ", res.Remaincredit)
		fmt.Println("Type 	= ", res.Type)
	}
}
