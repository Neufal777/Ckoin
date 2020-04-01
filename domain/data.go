package domain

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func Pricebtc(w http.ResponseWriter, r *http.Request) {
	resp, err := http.Get("https://www.bitstamp.net/api/ticker/")

	if err != nil {
		log.Println(err)
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)

	defer resp.Body.Close()
	//jsonData, err := json.MarshalIndent(body, "", "    ")

	res := string(body)
	fmt.Fprintf(w, res)
}
