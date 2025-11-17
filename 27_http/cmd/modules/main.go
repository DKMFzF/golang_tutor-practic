package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

func mainPage(res http.ResponseWriter, req *http.Request) {
	res.WriteHeader(http.StatusAccepted)
	res.Write([]byte("Привет!"))
}

func apiRoute(res http.ResponseWriter, req *http.Request) {
	body := fmt.Sprintf("Method: %s\r\n", req.Method)
	body += "Header =============\r\n"
	for k, n := range req.Header {
		body += fmt.Sprintf("%s: %v\r\n", k, n)
	}

	body += "Query parameters ============\r\n"
	//for k, v := range req.URL.Query() {
	//body += fmt.Sprintf("%s: %v\r\n", k, v)
	//}

	if err := req.ParseForm(); err != nil {
		res.Write([]byte(err.Error()))
		return
	}

	for k, v := range req.Form {
		body += fmt.Sprintf("%s: %v\r\n", k, v)
	}

	res.WriteHeader(http.StatusAccepted)
	res.Write([]byte(body))
}

func postReq(res http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodPost {
		http.Error(res, "Плохой метод", http.StatusBadRequest)
		return
	}

	body, err := io.ReadAll(req.Body)
	if err != nil {
		http.Error(res, "Нет body", http.StatusBadRequest)
		return
	}

	res.WriteHeader(http.StatusAccepted)
	res.Write([]byte(body))
}

type Subj struct {
	Product string `json:"product"`
	Price   int    `json:"price"`
}

func jsonHandler(w http.ResponseWriter, r *http.Request) {
	subj := Subj{"Mikl", 50}

	resp, err := json.Marshal(subj)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(resp)
}

func allMethod(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "1")
	fmt.Fprint(w, "2")
	w.Write([]byte("3"))
}

// middlewares

func middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%v", r.Method)
		next.ServeHTTP(w, r)
	})
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc(`/api/`, apiRoute)
	mux.Handle(`/`, middleware(http.HandlerFunc(mainPage)))
	mux.HandleFunc(`/post-test`, postReq)
	mux.HandleFunc(`/json`, jsonHandler)
	mux.HandleFunc(`/all`, allMethod)

	log.Print("Start server...")
	if err := http.ListenAndServe(`:8080`, mux); err != nil && err != http.ErrServerClosed {
		panic(err)
	}
}
