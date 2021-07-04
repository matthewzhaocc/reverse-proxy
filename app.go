package main

import (
	"net/http"
	"net/http/httputil"
	"net/url"
	"github.com/gorilla/mux"
	"encoding/json"
)

func init() {
	file, err :=
}

func main() {
	mux := mux.NewRouter()
	rProxyBackend := httputil.NewSingleHostReverseProxy(&url.URL{
		Scheme: "https",
		Host: "aws.amazon.com",
	})
	route := mux.Host("sus.local")
	route.HandlerFunc(rProxyBackend.ServeHTTP)
	
	http.ListenAndServe(":3000", mux)
	
}