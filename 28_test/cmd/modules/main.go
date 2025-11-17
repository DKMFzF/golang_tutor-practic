package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
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

func redirect(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, `/api`, http.StatusMovedPermanently)
}

func main() {
	log.Print("Start server...")

	if err := run(); !errors.Is(err, http.ErrServerClosed) {
		os.Exit(1)
	}
}

func run() error {
	srv := &http.Server{
		Addr:    ":8080",
		Handler: handlers(),
	}
	return srv.ListenAndServe()
}

func handlers() *http.ServeMux {
	mux := http.NewServeMux()

	staticDir := http.FileServer(http.Dir(`./static`))
	mux.Handle(`GET /static/`, http.StripPrefix(`/static/`, staticDir))

	mux.Handle(`GET /api/`, http.HandlerFunc(apiRoute))
	mux.Handle(`/`, middleware(http.HandlerFunc(mainPage)))
	mux.Handle(`POST /post-test`, http.HandlerFunc(postReq))
	mux.Handle(`GET /json`, http.HandlerFunc(jsonHandler))
	mux.Handle(`GET /all`, http.HandlerFunc(allMethod))
	mux.Handle(`GET /redirect`, http.HandlerFunc(redirect))
	mux.Handle(`GET /redirect-std`, http.RedirectHandler(`/api`, http.StatusMovedPermanently))

	return mux
}

func Sum(v ...int) int {
	var sum int
	for _, v := range v {
		sum += v
	}

	return sum
}
