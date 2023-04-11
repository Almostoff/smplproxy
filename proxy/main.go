package main

import (
	"bytes"
	"io/ioutil"
	"log"
	"net/http"
)

var (
	counter = 0
	fH      = "http://localhost:8181"
	sH      = "http://localhost:8282"
	adr     = "http://localhost:"
)

func main() {
	http.HandleFunc("/", handler)
	log.Print("proxy starting server")
	log.Fatalln(http.ListenAndServe("localhost:8080", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	if counter == 0 {
		resp := buildRequest(r, "8181")
		buildResponse(w, 200, resp)
		counter++
		return
	}

	resp := buildRequest(r, "8182")
	buildResponse(w, 200, resp)
	counter--
}

func buildResponse(w http.ResponseWriter, statusCode int, body []byte) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(statusCode)
	w.Write(body)
}

func buildRequest(r *http.Request, port string) []byte {
	req, err := http.NewRequest(r.Method, adr+port+r.URL.Path, bytes.NewBuffer([]byte(nil)))
	if err != nil {
		log.Fatal(err)
	}
	req.Header = r.Header
	req.Body = r.Body
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	respBody, _ := ioutil.ReadAll(resp.Body)
	return respBody
}
