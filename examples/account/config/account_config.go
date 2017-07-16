package main

import (
	"fmt"
	"net/url"

	"github.com/kavenegar/kavenegar-go"
	"github.com/kavenegar/kavenegar-go/enums/account/apilog"
	"github.com/kavenegar/kavenegar-go/enums/account/dailyreport"
	"github.com/kavenegar/kavenegar-go/enums/account/debugmode"
	"github.com/kavenegar/kavenegar-go/enums/account/resendfailed"
)

func main() {
	api := kavenegar.NewKavenegar(" your apikey ")
	params := &kavenegar.AccountConfigParam{
		Apilogs:     AccountAPILog.Enabled,
		Dailyreport: AccountDailyReport.Enabled,
		Debugmode:   AccountDebugMode.Enabled,
		//Defaultsender : "",
		//Mincreditalarm :" ",
		Resendfailed: AccountResendFailed.Enabled,
	}

	if res, err := api.Account.Config(params); err != nil {
		switch err := err.(type) {
		case *kavenegar.APIError:
			fmt.Println(err.Error())
		case *kavenegar.HTTPError:
			fmt.Println(err.Error())
		default:
			fmt.Println(err.Error())
		}
	} else {
		fmt.Println("Apilogs 	= ", res.Apilogs)
		fmt.Println("Dailyreport 	= ", res.Dailyreport)
		fmt.Println("Debugmode 	= ", res.Debugmode)
		fmt.Println("Defaultsender 	= ", res.Defaultsender)
		fmt.Println("Mincreditalarm 	= ", res.Mincreditalarm)
		fmt.Println("Resendfailed 	= ", res.Resendfailed)
		//...

	}

	//Account.CreateAccountConfig
	v := url.Values{}
	v.Set("apilogs", AccountAPILog.Enabled.String())
	v.Set("dailyreport", AccountAPILog.Enabled.String())
	v.Set("debugmode", AccountAPILog.Enabled.String())
	//v.Set("defaultsender", "")
	//v.Set("mincreditalarm", "")
	v.Set("resendfailed", AccountAPILog.Enabled.String())
	if res, err := api.Account.CreateConfig(v); err != nil {
		switch err := err.(type) {
		case *kavenegar.APIError:
			fmt.Println(err.Error())
		case *kavenegar.HTTPError:
			fmt.Println(err.Error())
		default:
			fmt.Println(err.Error())
		}
	} else {
		fmt.Println("Apilogs 	= ", res.Apilogs)
		fmt.Println("Dailyreport 	= ", res.Dailyreport)
		fmt.Println("Debugmode 	= ", res.Debugmode)
		fmt.Println("Defaultsender 	= ", res.Defaultsender)
		fmt.Println("Mincreditalarm 	= ", res.Mincreditalarm)
		fmt.Println("Resendfailed 	= ", res.Resendfailed)
		//...
	}

}
