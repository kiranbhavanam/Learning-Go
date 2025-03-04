package ch13

import (
	"fmt"
	"net/http"
)

// type welcome string

// func (wc welcome) ServeHTTP(w http.ResponseWriter,r *http.Request){
// 	fmt.Fprintf(w,"Welcome to our server")
// }
func Server2() {

	router:=http.NewServeMux()
	var wc welcome
	router.Handle("/",wc)
	router.HandleFunc("/login",func(w http.ResponseWriter,r *http.Request){
		fmt.Fprintf(w,"login")
	})
	router.HandleFunc("/logout",func(w http.ResponseWriter,r *http.Request){
		fmt.Fprintf(w,"logout")
	})
	server := http.Server{
		Addr:":8080",
		Handler: router,
	}
	server.ListenAndServe()
}