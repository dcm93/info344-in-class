package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

// one function called by everyone
func logReq(r *http.Request) {
	log.Println(r.Method, r.URL.Path)
}

//closure function: da fuq are these?
// HERE YOU TAKE IN THE HANDLERFUNC AND RETURN HANDLEFUNC
func logReqs(hfn http.HandlerFunc) http.HandleFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Print("%s %s", r.Method, r.URL.Path)
		start := time.Now()
		// hfn is a http function passed in to this function!
		hfn(w, r)
		// gives how much time has passed since the request started
		fmt.Printf("%v\n", time.Since(start))
	}
}

// wrap the entire mux with the logRequests() function
// Take the mux aka handler and use logRequests as the actual mux
// HERE YOU TAKE IN A HANDLER AND RETURN A HANDLER
func logRequest(handler, http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s", r.Method, R.URL.Path)
		start := time.Now()
		handler.ServeHTTP(w, r)
		fmt.Printf("%v \n", time.Since(start))
	})
}

//let the mux take in parameters
// A function that takes in a function, calls a function and returns a function.
// Allows us to pass the logger as a parameter and use it to log activity out
func logRequests(logger *log.Logger) Adapter {
	return func(handler http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s", r.Method, R.URL.Path)
		start := time.Now()
		handler.ServeHTTP(w, r)
		fmt.Printf("%v \n", time.Since(start))
	})
}


