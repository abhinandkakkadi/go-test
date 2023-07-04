package main

import (
	"net/http"
	"strings"
	"testing"

	"github.com/go-chi/chi"
)

// app in the middle to signify the receiver (not mandatory- just for identification purpose)
func Test_application_routes(t *testing.T) {

	var registered = []struct{
		route string
		method string

	}{
		{"/","GET"},
		{"/static/*","GET"},
	}

	var app application
	mux := app.routes()

	// chiRoutes will have details of all routes registered with the handler
	chiRoutes := mux.(chi.Routes)

	for _, route := range registered {
		// check to see if the route exist
		if !routeExists(route.route, route.method,chiRoutes) {
			t.Errorf("route %s is not registered",route.route)
		}
	}

}

// true if the route exists and false if the route does not exists 
func routeExists(testRoute,testMethod string,chiRoutes chi.Routes) bool {

	found :=  false

	// it will walk over each routes and it's associated method that is registered ti the chi router
	// and check if the testRoute and it's correspond testMethod match along with any of the given routes
	_ = chi.Walk(chiRoutes, func(method string, route string, handler http.Handler, middleware ...func(http.Handler) http.Handler) error {

		if  strings.EqualFold(method,testMethod) && strings.EqualFold(route, testRoute) {
			found = true
		}

		return nil
	})

	return found

}