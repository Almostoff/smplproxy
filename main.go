package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"log"
	"net/http"
)

type PostData struct {
	Data string `json:"data"`
}

type PatchData struct {
	Data string `json:"data"`
}

type PutData struct {
	Data string `json:"data"`
}

var adr = "http://localhost"

const port = 8181

func main() {
	port := flag.Int("port", port, "port to listen on")
	flag.Parse()

	adr = fmt.Sprintf(":%d", *port)
	log.Printf("Starting server on %s for testing HTTP...\n", adr)

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Route("/test", func(r chi.Router) {
		r.Post("/", postH)
		r.Patch("/", patchH)
		r.Put("/", putH)

		r.Route("/{data}", func(r chi.Router) {
			r.Delete("/", deleteH)
			r.Get("/", getH)
		})
	})

	if err := http.ListenAndServe(adr, r); err != nil {
		log.Fatal(err)
	}
}

func getH(w http.ResponseWriter, r *http.Request) {
	log.Printf("Handler Get request on %s port", adr)

	data := chi.URLParam(r, "data")

	result := map[string]string{
		"data":            data,
		"service_on_port": adr,
	}

	response, err := json.Marshal(result)
	if err != nil {
		buildResponse(w, http.StatusInternalServerError, nil)
		return
	}
	buildResponse(w, 200, response)
}

func deleteH(w http.ResponseWriter, r *http.Request) {
	log.Printf("Handler Delete request on %s port", adr)

	deletedata := chi.URLParam(r, "data")

	result := map[string]string{
		"data":            deletedata,
		"service_on_port": adr,
	}

	response, err := json.Marshal(result)
	if err != nil {
		buildResponse(w, http.StatusInternalServerError, nil)
		return
	}
	buildResponse(w, http.StatusOK, response)
}

func putH(w http.ResponseWriter, r *http.Request) {
	log.Printf("Handler Put request on %s port", adr)

	putdata := &PutData{}
	err := json.NewDecoder(r.Body).Decode(putdata)
	if err != nil {
		buildResponse(w, http.StatusInternalServerError, nil)
		return
	}

	result := map[string]string{
		"data":            putdata.Data,
		"service_on_port": adr,
	}

	response, err := json.Marshal(result)
	if err != nil {
		buildResponse(w, http.StatusInternalServerError, nil)
		return
	}
	buildResponse(w, 200, response)
}

func postH(w http.ResponseWriter, r *http.Request) {
	log.Printf("Handler Post request on %s port", adr)

	postdata := &PostData{}
	err := json.NewDecoder(r.Body).Decode(postdata)
	if err != nil {
		buildResponse(w, http.StatusInternalServerError, nil)
		return
	}

	result := map[string]string{
		"data":            postdata.Data,
		"service_on_port": adr,
	}

	response, err := json.Marshal(result)
	if err != nil {
		buildResponse(w, http.StatusInternalServerError, nil)
		return
	}
	buildResponse(w, http.StatusCreated, response)
}

func patchH(w http.ResponseWriter, r *http.Request) {
	log.Printf("Handler Patch request on %s port", adr)

	patchdata := &PatchData{}
	err := json.NewDecoder(r.Body).Decode(patchdata)
	if err != nil {
		buildResponse(w, http.StatusInternalServerError, nil)
		return
	}

	result := map[string]string{
		"data":            patchdata.Data,
		"service_on_port": adr,
	}

	response, err := json.Marshal(result)
	if err != nil {
		buildResponse(w, http.StatusInternalServerError, nil)
		return
	}
	buildResponse(w, 200, response)
}

func buildResponse(w http.ResponseWriter, statusCode int, body []byte) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(statusCode)
	w.Write(body)
}
