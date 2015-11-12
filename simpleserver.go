package main

import (
	"github.com/julienschmidt/httprouter"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
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
	hostAndPort := "localhost:3000"
	log.Printf("Running on %v", hostAndPort)
	http.ListenAndServe(hostAndPort, router)
}
