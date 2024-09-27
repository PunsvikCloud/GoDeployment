package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, err := w.Write([]byte(getStringWithTimestamp()))
		if err != nil {
			return
		}
	})

	log.Fatal(http.ListenAndServe(":8080", nil)) //nolint:gosec
}

// getStringWithTimestamp returns the string with "OK" and the current timestamp
func getStringWithTimestamp() string {
	return fmt.Sprintf("OK - %s", time.Now().Format(time.RFC3339))
}
