package ch13

import (
	"fmt"
	"net/http"
)

type welcome string

func (wc welcome) ServeHTTP(w http.ResponseWriter,r *http.Request){
	// x:=10
	fmt.Fprintf(w,"Welcome to our server")
}
func Server() {


	var wc welcome
	server := http.Server{
		Addr:":8080",
		Handler: wc,
	}
	server.ListenAndServe()
}