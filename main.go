package main

import (
	"fmt"
	"html"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func main() {
	//SimpleHTTP("http://example.com")
	//ClientHTTP()
	//HeaderHTTP()
	//TransportHTTP()
	SimpleWebServer()
}

func SimpleHTTP(u string) {
	res, err := http.Get(u)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("%s", body)
}

func ClientHTTP() {
	client := &http.Client{}
	res, err := client.Get("http://example.com")
	if err != nil {
		log.Fatal("ERROR: could not reach http://example.com")
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	log.Printf("%s", body)
}

func HeaderHTTP() {
	client := &http.Client{}
	req, err := http.NewRequest("GET", "http://example.com", nil)
	if err != nil {
		log.Fatal("ERROR: could not initialize request")
	}
	req.Header.Add("If-None-Match", `W/"wyzzy"`)
	res, err := client.Do(req)
	if err != nil {
		log.Fatal("ERROR: could not reach http://example.com")
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	log.Printf("%s", body)
}

func TransportHTTP() {
	tr := &http.Transport{
		MaxIdleConns:       10,
		IdleConnTimeout:    30 * time.Second,
		DisableCompression: true,
	}
	client := &http.Client{Transport: tr}
	req, err := http.NewRequest("GET", "http://example.com", nil)
	if err != nil {
		log.Fatal("ERROR: could not initialize request")
	}
	req.Header.Add("If-None-Match", `W/"wyzzy"`)
	res, err := client.Do(req)
	if err != nil {
		log.Fatal("ERROR: could not reach http://example.com")
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	log.Printf("%s", body)
}

func SimpleWebServer() {
	http.HandleFunc("/foo", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
	})

	log.Fatal(http.ListenAndServe(":8000", nil))
}
