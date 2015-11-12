package main

import (
	"flag"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

var (
	defaultBindAddress = "127.0.0.1"
	defaultBindPort    = 3000
)

func main() {
	var bindAddress string
	var bindPort int

	flag.StringVar(&bindAddress, "address", defaultBindAddress, "the address to bind the server to")
	flag.IntVar(&bindPort, "port", defaultBindPort, "the port to bind the server to")

	hostAndPort := net.JoinHostPort(bindAddress, strconv.Itoa(bindPort))
	log.Printf("Running on %v", hostAndPort)
	err := http.ListenAndServe(hostAndPort, router())
	if err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}

func router() http.Handler {
	router := httprouter.New()
	router.GET("/", func(res http.ResponseWriter, req *http.Request, params httprouter.Params) {
		log.Printf("GET /")
	})
	router.POST("/", func(res http.ResponseWriter, req *http.Request, params httprouter.Params) {
		data, err := ioutil.ReadAll(req.Body)
		if err != nil {
			log.Printf("Error: %v", err)
		} else {
			log.Printf("Data: %v", string(data))
		}
	})
	router.PUT("/", func(res http.ResponseWriter, req *http.Request, params httprouter.Params) {
		data, err := ioutil.ReadAll(req.Body)
		if err != nil {
			log.Printf("Error: %v", err)
		} else {
			log.Printf("Data: %v", string(data))
		}
	})
	router.DELETE("/", func(res http.ResponseWriter, req *http.Request, params httprouter.Params) {
		log.Printf("DELETE /")
	})

	return router
}
