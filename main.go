package main

import (
	"log"
	"net/http"

	app "github.com/ckoin/app"
)

func main() {

	info := app.PriceCoin{
		High:      6511.60000000,
		Last:      6198.04,
		Timestamp: 1585760373,
		Bid:       6191.84,
		Vwap:      6312.49,
		Volume:    6685.93951412,
		Low:       6137.71000000,
		Ask:       6198.04,
		Open:      6428.74,
	}

	log.Println(info)
	http.HandleFunc("/", app.Pricebtc)
	http.ListenAndServe(":8000", nil)
}
