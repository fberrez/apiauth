package main

import (
	"log"
	"net/http"

	"github.com/fberrez/apiauth/api"
)

func main() {
	api := api.New("mysecret", 1000)
	log.Println("server started and listening on http://127.0.0.1:8080")
	http.ListenAndServe("127.0.0.1:8080", api)
}
