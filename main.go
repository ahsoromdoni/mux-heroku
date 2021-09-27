package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func main()  {
	r := mux.NewRouter()
	r.HandleFunc("/handler", handler).Methods("GET")
	http.Handle("/",r)
	http.ListenAndServe(nil, nil)
}

func handler(w http.ResponseWriter, r *http.Request)  {
	fmt.Fprint(w, "hello")
}
