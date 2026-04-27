package handlers

import (
	"html/template"
	"net/http"
	"fmt"
)

func HomeHandler(w http.ResponseWriter, req *http.Request) {
	tmpl, err := template.ParseFiles("./static/pages/index.html")
	if err != nil {
		fmt.Fprintf(w, "Found error %v", err)
		return
	}
	tmpl.Execute(w, "index.html")
}