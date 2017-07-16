[![Build Status](https://travis-ci.org/KaveNegar/kavenegar-go.svg?branch=master)](https://travis-ci.org/KaveNegar/kavenegar-go)
 
# kavenegar-go 
 
 
## Installation
```
go get github.com/kavenegar/kavenegar-go
```
## Usage
### Send
```golang
package main
import (
	"fmt"
	"github.com/kavenegar/kavenegar-go"
)
func main() {
	sender := ""                 //Sender Line Number(optional)
	receptor := []string{"", ""} //Recipient numbers
	message := "Hello Go!"       //Text message
	params := &kavenegar.MessageParam{}
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
}
```
### OTP
```golang
package main
import (
	"fmt"
	"github.com/kavenegar/kavenegar-go"
)
func main() {
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
}
```
## Contribution
