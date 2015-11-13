package main

import (
	"flag"
	"fmt"
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
		log.Printf("POST /")
		echo(res, req, http.StatusOK)
	})
	router.PUT("/", func(res http.ResponseWriter, req *http.Request, params httprouter.Params) {
		log.Printf("PUT /")
		echo(res, req, http.StatusOK)
	})
	router.DELETE("/", func(res http.ResponseWriter, req *http.Request, params httprouter.Params) {
		log.Printf("DELETE /")
		res.WriteHeader(http.StatusNoContent)
	})

	router.GET("/:code", func(res http.ResponseWriter, req *http.Request, params httprouter.Params) {
		log.Printf("GET /%v", params.ByName("code"))
		code, err := strconv.Atoi(params.ByName("code"))
		if err != nil {
			http.Error(res, fmt.Sprintf("Error converting requested status code to int: %v", err), http.StatusBadRequest)
		} else {
			res.WriteHeader(code)
		}
	})
	router.POST("/:code", func(res http.ResponseWriter, req *http.Request, params httprouter.Params) {
		log.Printf("POST /%v", params.ByName("code"))
		code, err := strconv.Atoi(params.ByName("code"))
		if err != nil {
			http.Error(res, fmt.Sprintf("Error converting requested status code to int: %v", err), http.StatusBadRequest)
		} else {
			echo(res, req, code)
		}
	})
	router.PUT("/:code", func(res http.ResponseWriter, req *http.Request, params httprouter.Params) {
		log.Printf("PUT /%v", params.ByName("code"))
		code, err := strconv.Atoi(params.ByName("code"))
		if err != nil {
			http.Error(res, fmt.Sprintf("Error converting requested status code to int: %v", err), http.StatusBadRequest)
		} else {
			echo(res, req, code)
		}
	})
	router.DELETE("/:code", func(res http.ResponseWriter, req *http.Request, params httprouter.Params) {
		log.Printf("DELETE /%v", params.ByName("code"))
		code, err := strconv.Atoi(params.ByName("code"))
		if err != nil {
			http.Error(res, fmt.Sprintf("Error converting requested status code to int: %v", err), http.StatusBadRequest)
		} else {
			res.WriteHeader(code)
		}
	})

	return router
}

func echo(res http.ResponseWriter, req *http.Request, code int) {
	defer req.Body.Close()
	data, err := ioutil.ReadAll(req.Body)
	if err != nil {
		log.Printf("Error: %v", err)
	} else {
		log.Printf("Data: %v", string(data))
	}
	res.Header().Add("Content-type", req.Header.Get("Content-type"))
	res.WriteHeader(code)
	res.Write(data)
}
