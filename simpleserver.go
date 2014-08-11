package main

import (
	"github.com/go-martini/martini"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	m := martini.Classic()
	m.Get("/", func(res http.ResponseWriter, req *http.Request) string {
		log.Printf("GET /")
		return "GET"
	})
	m.Post("/", func(res http.ResponseWriter, req *http.Request) string {
		data, err := ioutil.ReadAll(req.Body)
		if err != nil {
			log.Printf("Error: %v", err)
		} else {
			log.Printf("Data: %v", string(data))
		}

		return "POST"
	})
	m.Put("/", func(res http.ResponseWriter, req *http.Request) string {
		data, err := ioutil.ReadAll(req.Body)
		if err != nil {
			log.Printf("Error: %v", err)
		} else {
			log.Printf("Data: %v", string(data))
		}

		return "PUT"
	})
	m.Delete("/", func(res http.ResponseWriter, req *http.Request) string {
		log.Printf("DELETE /")
		return "DELETE"
	})
	hostAndPort := "localhost:4000"
	log.Printf("Running on %v", hostAndPort)
	http.ListenAndServe(hostAndPort, m)
}
