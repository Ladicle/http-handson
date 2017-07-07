package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strings"
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
	itemKeys := []string{}
	if v := r.FormValue("item"); v != "" {
		itemKeys = append(itemKeys, v)
	}
	if cookie, err := r.Cookie("cart"); err == nil {
		itemKeys = append(itemKeys, strings.Split(cookie.Value, ",")...)
	}

	cart := Cart{ItemList: []Item{}}
	for _, itemKey := range itemKeys {
		if item, ok := catalog.ItemMap[itemKey]; ok {
			cart.addItem(item)
		}
	}

	http.SetCookie(w, &http.Cookie{Name: "cart", Value: strings.Join(itemKeys, ",")})
	generateTemplate(w, "cart", cart)
}

func purchaseHandler(w http.ResponseWriter, r *http.Request) {
	http.SetCookie(w, &http.Cookie{Name: "cart", Value: ""})
	generateTemplate(w, "purchase", nil)
}
