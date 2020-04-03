package domain

import (
	"fmt"
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

func Pricebtc(w http.ResponseWriter, r *http.Request) {

	example := &PriceCoin{
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

	// resp, err := http.Get("https://www.bitstamp.net/api/ticker/")

	// if err != nil {
	// 	log.Println(err)
	// }
	// defer resp.Body.Close()
	// body, _ := ioutil.ReadAll(resp.Body)

	// defer resp.Body.Close()
	// //jsonData, err := json.MarshalIndent(body, "", "    ")

	// res := string(body)

	i := example.Last
	x := fmt.Sprintf("%f", i)
	fmt.Fprintf(w, x)
}
