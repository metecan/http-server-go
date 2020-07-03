package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Method:", r.Method, "\nStatus:", http.StatusOK)
	})
	http.ListenAndServe(":80", nil)
}