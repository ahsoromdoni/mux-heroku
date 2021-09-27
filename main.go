package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
<<<<<<< HEAD
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
=======
	r.HandleFunc("/handler", handler).Methods("GET")
	http.Handle("/",r)
	http.ListenAndServe(nil, nil)
>>>>>>> e9a6f2c6e4f3188f0a0717d82be3de36238a6a8a
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "hello")
}
