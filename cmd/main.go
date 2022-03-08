package main

import (
	"fmt"
	"log"
	"os"

	"github.com/tgirier/envelope"
)

func main() {
	var port = os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Fatal(envelope.ListenAndServe(fmt.Sprintf(":%s", port)))
}
