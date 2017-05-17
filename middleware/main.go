package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

//Starts a web server listening in 4000 and add 3 handlers.
// do a go build & install with this main
// handlers, what they do is the serve one particular url

// One shared function:
// once you add the logging function and add a call to logReq in handler function, you get logging written to the server
// Not a great idea because once the request is completed, no way to check how long it took or learn more.

//Closure function:
//to use a closure, you wrap HelloHandler1 with logReqs()

//Wrapper function:
//You give it as a param to http.ListenAndServe which calls logRequests, which in turn calls the MUX, which in turn calls the handler.
func main() {
	addr := "localhost:4000"

	mux := http.NewServeMux()
	muxLogged := http.NewServeMux()

	// exchange mux for http if you dont do entire wrapper.
	muxLogged.HandleFunc("/v1/hello1", logReqs(HelloHandler1))
	muxLogged.HandleFunc("/v1/hello2", HelloHandler2)
	//mux.HandleFunc("/v1/hello3", HelloHandler3)

	//These two lines makes it that urls served by HelloHandler3 do not log anything out.
	// handle() to add one mux to another mux
	mux.HandleFunc("/v1/hello3", HelloHandler3)
	logger := log.New(os.Stdout, "", log.LstdFlags)
	//returns a function with the first call and then that function gets immediately executed with the mux you give it.
	//mux.Handle("/v1/", logRequests(logger)(muxLogged))

	// using an adapter to use various adapters for 1 handler
	mux.Handle("/v1/", Adapt(muxLogged, logRequests(logger), throttleRequests(2, time.Minute)))

	fmt.Printf("listening at %s...\n", addr)
	log.Fatal(http.ListenAndServe(addr, logRequests(mux)))
	log.Fatal(http.ListenAndServe(addr, nil))
}
