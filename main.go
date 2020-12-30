package main

import (
	"fmt"
	"log"
	"net/http"

	f "./controller"
	"./repository"
)

func main() {
	mux := f.Ping()
	db := repository.Connect()
	defer db.Close()
	fmt.Println("Serving request..")
	log.Fatal(http.ListenAndServe("localhost:3000", mux))
}
