package main

import (
	"net/http"

	domain "github.com/ckoin/domain"
)

func main() {

	http.HandleFunc("/", domain.Pricebtc)
	http.ListenAndServe(":8000", nil)
}
