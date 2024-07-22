package main

import (
	"FirstProject/pkg/api"
	"fmt"
	"log"
	"net/http"
)

func main() {
	api.FillEndpoints()

	fmt.Println("Server is working...")

	log.Fatal(http.ListenAndServe("localhost:8090", nil))
}
