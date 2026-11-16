package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type PageData struct {
	Result string
}

var tmpl *template.Template

func homeHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	err := tmpl.Execute(w, PageData{})
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}

func asciiArtHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	// err := r.ParseForm()
	// if err != nil {
	// 	http.Error(w, "Bad Request", http.StatusBadRequest)
	// 	return
	// }
	text := r.FormValue("text")
	banner := r.FormValue("banner")

	if err := validateInput(text); err != nil {
		http.Error(w, "Bad Request: "+err.Error(), http.StatusBadRequest)
		return
	}

	if banner != "standard" && banner != "shadow" && banner != "thinkertoy" {
		http.Error(w, "Bad Request: invalid banner selection", http.StatusBadRequest)
		return
	}
	bannerFile := banner + ".txt"

	rawBanner, err := readFile(bannerFile)
	if err != nil {
		http.NotFound(w, r)
		return
	}

	charMap, err := buildBannerMap(rawBanner)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	lines := slashNHandling(text)
	result := renderArt(charMap, lines)

	err = tmpl.Execute(w, PageData{Result: result})
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}

func main() {
	var err error
	tmpl, err = template.ParseFiles("templates/index.html")
	if err != nil {
		fmt.Println("Error: could not parse template —", err)
		return
	}

	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/ascii-art", asciiArtHandler)

	fmt.Println("Server starting on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
