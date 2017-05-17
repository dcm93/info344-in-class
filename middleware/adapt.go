package main

import (
	"net/http"
)

//Adapter is the type for adapter functions.
//An adapter function accepts an http.Handler
//and returns a new http.Handler that wraps the
//input handler, providing some pre- and/or
//post-processing.
type Adapter func(http.Handler) http.Handler

//TODO: write an Adapt() function that accepts:
// - handler http.Handler the handler to adapt
// - a variadic slice of Adapter functions
//iterate the slice of Adapter functions in
//reverse order, passing the `handler` to
//each, and resetting `handler` to the
//handler returned from the Adapter func

//takes any number of adapters, wraps the handler with all of them and returns it
// allows us to not have to do the (1)(2) call in main.go
func Adapt(handler http.Handler, adapters ...Adapter) http.Handler{
	for idx = len(adapters) -1; idx >= 0; idx--{
		// gives in adapter function that will take in the handler
		handler = adapters[idx](handler)
	}
	return handler
}