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
	"net/url"
	"github.com/kavenegar/kavenegar-go"
)
func main() {
	api := kavenegar.NewKavenegar(" your apikey ")
	sender := ""                 
	receptor := []string{"", ""}
	message := "Hello Go!" 
	if res, err := api.Message.Send(sender, receptor, message, nil); err != nil {
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
	"net/url"
	"github.com/kavenegar/kavenegar-go"
)
func main() {
	api := kavenegar.NewKavenegar(" your apikey ")
	receptor := ""
	template := ""
	token := ""
	params := &kavenegar.VerifyLookupParam{
	}
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
### Send Bulk
```golang
package main
import (
	"fmt"
	"net/url"
	"github.com/kavenegar/kavenegar-go"
)
func main() {
	api := kavenegar.New(" your apikey here ")	
	res, err := api.Message.SendArray(url.Values{
		"receptor": {"",""},
		"message": {"Hello Go!","Hello Go!"},
		"sender": {"",""},
	})
	if err != nil {
         switch err := err.(type) {
			case *kavenegar.APIError:
				fmt.Println(err.Error())
			case *kavenegar.HTTPError:
				fmt.Println(err.Error())
			default:
				fmt.Println(err.Error())
         }
	}else{
		fmt.Println(res)
	}
}
```
## Contribution
