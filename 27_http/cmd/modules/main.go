package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func mainPage(res http.ResponseWriter, req *http.Request) {
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

	res.Write([]byte(body))
}

func main() {
	mux := http.NewServeMux()

	go mux.HandleFunc(`/api/`, apiRoute)
	go mux.HandleFunc(`/`, mainPage)
	go mux.HandleFunc(`/post-test`, postReq)

	log.Print("Start server...")
	if err := http.ListenAndServe(`:8080`, mux); err != nil && err != http.ErrServerClosed {
		panic(err)
	}
}
