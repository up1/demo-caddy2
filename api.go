package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/api/test", func(w http.ResponseWriter, r *http.Request) { fmt.Fprintf(w, "test") })
	http.HandleFunc("/api/time", func(w http.ResponseWriter, r *http.Request) { fmt.Fprintf(w, time.Now().Format("02 Jan 2006")) })
	http.ListenAndServe(":8000", nil)

}
