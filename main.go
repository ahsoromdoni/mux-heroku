package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/ping", handler).Methods("GET")
	http.Handle("/", r)
	address := resolveAddress([]string{})
	http.ListenAndServe(address, nil)
}

func resolveAddress(addr []string) string {
	switch len(addr) {
	case 0:
		if port := os.Getenv("PORT"); port != "" {
			// debugPrint("Environment variable PORT=\"%s\"", port)
			return ":" + port
		}
		// debugPrint("Environment variable PORT is undefined. Using port :8080 by default")
		return ":8080"
	case 1:
		return addr[0]
	default:
		panic("too many parameters")
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "hello")
}
