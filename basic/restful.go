package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Payload struct {
	Stuff Data
}

type Data struct {
	Fruit  Fruits
	Vegies Vegetables
}

type Fruits map[string]int
type Vegetables map[string]int

func handlerICon(w http.ResponseWriter, r *http.Request) {
	fmt.Println("handle favicon call from web browser")
}

func serveRest(w http.ResponseWriter, r *http.Request) {
	fmt.Println("serveRest")
	respone, err := getJsonResponse()
	// fmt.Println("json is :", respone)
	if err != nil {
		panic(err)
	}
	fmt.Fprintf(w, string(respone))
}

func getJsonResponse() ([]byte, error) {
	fruits := make(map[string]int)
	fruits["apple"] = 1
	fruits["orange"] = 2

	vegetables := make(map[string]int)
	vegetables["carrot"] = 11
	vegetables["peppers"] = 22

	d := Data{fruits, vegetables}
	p := Payload{d}
	return json.MarshalIndent(p, "", "  ")
}

func main() {
	http.HandleFunc("/favicon.ico", handlerICon)
	http.HandleFunc("/", serveRest)
	http.ListenAndServe("localhost:1337", nil)
}
