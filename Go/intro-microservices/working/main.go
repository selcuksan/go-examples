package main

import (
	"log"
	"microservices/working/handlers"
	"net/http"
	"os"
)

func main() {
	l := log.New(os.Stdout, "product-api", log.LstdFlags)

	hh := handlers.NewHello(l)

	sm := http.NewServeMux()

	sm.Handle("/", hh)
	log.Println("Starting Server")
	err := http.ListenAndServe(":9090", sm)
	log.Fatal(err)
}
