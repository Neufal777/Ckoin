package main

import (
	app "github.com/ckoin/app"
)

// Available := []string{
// 	"btcusd",
// 	"btceur",
// 	"eurusd",
// 	"xrpusd",
// 	"xrpeur",
// 	"xrpbtc",
// 	"ltcusd",
// 	"ltceur",
// 	"ltcbtc",
// 	"ethusd",
// 	"etheur",
// 	"ethbtc",
// 	"bchusd",
// 	"bcheur",
// 	"bchbtc",
// }

func main() {

	//Check prices of different cryptocurrencies (Bitstamp Price)
	crypt := []string{
		"btcusd",
		"btceur",
		"eurusd",
		"xrpusd",
		"xrpeur",
		"xrpbtc",
		"ltcusd",
		"ltceur",
		"ltcbtc",
		"ethusd",
		"etheur",
		"ethbtc",
		"bchusd",
		"bcheur",
		"bchbtc",
		//add as many as you want
	}

	
	app.CheckPrices(crypt)
}
