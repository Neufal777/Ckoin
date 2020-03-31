package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// type Estudios struct {
// 	Primaria    bool `json:"Primeria"`
// 	Secundaria  bool `json:"Secundaria"`
// 	Universidad bool `json:"Universidad"`
// }

// type Other struct {
// 	Bi      int32 `json:"Business_intelligence"`
// 	DataEng int32 `json:"Data_engineer"`
// }
// type Skills struct {
// 	Frontend int32 `json:"Frontend"`
// 	Backend  int32 `json:"Backend"`
// 	Other    Other `json:"Other_Skills"`
// }
// type Person struct {
// 	Name     string   `json:"Name"`
// 	Edad     int      `json:"Edad"`
// 	Estudios Estudios `json:"Estudios"`
// 	Skills   Skills   `json:"Skills"`
// 	sexo     struct {
// 		Mujer  int `json:"mujer"`
// 		Hombre int `json:"hombre"`
// 	}
// }

func main() {

	// naoufal := Person{
	// 	Name: "Naoufal dahouli",
	// 	Edad: 19,
	// 	Estudios: Estudios{
	// 		Primaria:    true,
	// 		Secundaria:  false,
	// 		Universidad: false,
	// 	},
	// 	Skills: Skills{
	// 		Frontend: 7,
	// 		Backend:  8,
	// 		Other: Other{
	// 			Bi:      2,
	// 			DataEng: 5,
	// 		},
	// 	},
	// }

	resp, err := http.Get("https://www.bitstamp.net/api/ticker/")

	if err != nil {
		log.Println(err)
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)

	defer resp.Body.Close()
	//jsonData, err := json.MarshalIndent(body, "", "    ")

	res := string(body)
	//jsonData, err := json.MarshalIndent(body, "", "    ")

	fmt.Println(res)
}
