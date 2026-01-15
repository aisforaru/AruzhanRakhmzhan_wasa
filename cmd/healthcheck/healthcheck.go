package main

import (
	"flag"
	"log"
	"net/http"
	"os"
	"strconv" // Importing strconv to convert port to string
)

func main() {
	var port = flag.Int("port", 3000, "HTTP port for healthcheck")
	flag.Parse()

	// Construct the healthcheck URL
	url := "http://localhost:" + strconv.Itoa(*port) + "/liveness"
	res, err := http.Get(url)
	if err != nil {
		log.Println("Error:", err) // Log the error message
		os.Exit(1)                 // Exit with a non-zero status code
	} else if res.StatusCode != http.StatusOK && res.StatusCode != http.StatusNoContent {
		_ = res.Body.Close()
		log.Println("Healthcheck request not OK:", res.Status) // Log healthcheck failure
		os.Exit(1)
	}

	// Close the response body and exit successfully
	_ = res.Body.Close()
	os.Exit(0)
}
