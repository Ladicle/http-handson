package main

import (
	"fmt"
	"html/template"
	"net/http"
)

// GenerateRouter bind handlers to a multiplexer
func GenerateRouter() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/", indexHandler)
	mux.HandleFunc("/catalog", catalogHandler)
	mux.HandleFunc("/cart", cartHandler)
	mux.HandleFunc("/purchase", purchaseHandler)
	return mux
}

var catalog = NewCatalog()

func generateTemplate(w http.ResponseWriter, tplName string, data interface{}) {
	tpl, err := template.ParseFiles(fmt.Sprintf("view/%s.html", tplName))
	if err != nil {
		w.WriteHeader(500)
		return
	}
	if err = tpl.Execute(w, data); err != nil {
		w.WriteHeader(500)
		return
	}
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	generateTemplate(w, "index", nil)
}

func catalogHandler(w http.ResponseWriter, r *http.Request) {
	generateTemplate(w, "catalog", catalog)
}

func cartHandler(w http.ResponseWriter, r *http.Request) {
}

func purchaseHandler(w http.ResponseWriter, r *http.Request) {
}
