package main

import (
	"net/http"

	app "github.com/ckoin/app"
)

func main() {

	http.HandleFunc("/", app.Home)
	http.ListenAndServe(":8000", nil)
}
