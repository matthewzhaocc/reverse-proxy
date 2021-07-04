package main

import (
	"net/http"
	"net/http/httputil"
	"net/url"
	"github.com/gorilla/mux"
	"encoding/json"
	"path/filepath"
	"io/ioutil"
)
var config ReverseProxyConfiguration
func init() {
	configFile, err := filepath.Abs("./config.json")
	if err != nil {
		panic(err.Error())
	}

	value, err := ioutil.ReadFile(configFile)
	if err != nil {
		panic(err.Error())
	}
	json.Unmarshal(value, &config)
}

func main() {
	muxRouter := mux.NewRouter()
	for _, frontend := range config.Frontends {
		route := muxRouter.Host(frontend.Hostname)
		for _, backend := range config.Hosts {
			if backend.FrontendHostName == frontend.Hostname {
				route.HandlerFunc(httputil.NewSingleHostReverseProxy(&url.URL{
					Scheme: backend.Scheme,
					Host: backend.Host,
				}).ServeHTTP)
			}
		}
	} 
	
	http.ListenAndServe(":3000", muxRouter)
	
}

type ReverseProxyConfiguration struct {
	Hosts []ReverseProxyHost `json:"hosts"`
	Frontends []Frontend `json:"frontends"`
}

type ReverseProxyHost struct {
	Scheme string `json:"scheme"`
	Host string	`json:"host"`
	FrontendHostName string `json:"frontend_hostname"`
}

type Frontend struct {
	Hostname string `json:"hostname"`
}