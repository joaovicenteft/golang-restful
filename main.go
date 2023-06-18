package main

import (
	"log"
	"net/http"
	"restfulImp/handlers"
)

func main() {
	http.HandleFunc("/users", handlers.HandleUsers)
	log.Fatal(http.ListenAndServe(":8081", nil))
}
