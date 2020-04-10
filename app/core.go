package app

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
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

		res := string(body)

		in := []byte(res)
		var raw map[string]interface{}
		if err := json.Unmarshal(in, &raw); err != nil {
			panic(err)
		}
		raw["count"] = 1
		out, _ := json.MarshalIndent(raw, "", "  ")

		c = strings.ToUpper(c)

		c1 := c[:3]
		c2 := c[3:6]

		c = c1 + " - " + c2

		println(c, "\n")
		println(string(out))
		println("\n")
	}

}
