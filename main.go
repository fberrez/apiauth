package main

import (
	"context"
	"log"
	"net/http"

	"github.com/fberrez/apiauth/api"
)

func main() {
	ctx := context.Background()
	api, err := api.New(ctx, "mysecret", 1000)
	if err != nil {
		panic(err)
	}
	log.Println("server started and listening on http://127.0.0.1:8080")
	http.ListenAndServe("127.0.0.1:8080", api)
}
