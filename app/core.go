package app

import (
	"io/ioutil"
	"log"
	"net/http"
)

type PriceCoin struct {
	High      float32
	Last      float32
	Timestamp int64
	Bid       float32
	Vwap      float32
	Volume    float32
	Low       float32
	Ask       float32
	Open      float32
}

func Home() {

	cryptos := []string{
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
	}

	for _, c := range cryptos {
		resp, err := http.Get("https://www.bitstamp.net/api/v2/ticker/" + c + "/")

		if err != nil {
			log.Println(err)
		}
		defer resp.Body.Close()
		body, _ := ioutil.ReadAll(resp.Body)

		defer resp.Body.Close()
		//jsonData, err := json.MarshalIndent(body, "", "    ")

		res := string(body)

		log.Println(c, " :", res)
	}

}
