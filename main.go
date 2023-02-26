package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/ratanpatil212/mongodb/router"
)

func main() {
	fmt.Println("MONGO DB API")
	http.Handle("/", http.FileServer(http.Dir("./static")))
	r := router.Router()
	fmt.Println("Server is getting started ...")
	log.Fatal(http.ListenAndServe(":4000", r))
	fmt.Println("Listening on port 4000")
}
