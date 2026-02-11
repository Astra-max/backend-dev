package handlers

import (
	"ascii-art-web/ascii-art"
	"html/template"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"encoding/json"
	"strings"
)

func renderTemplateWithData(w http.ResponseWriter, tmplName string, data interface{}) {
	lp := filepath.Join("templates", "layout.html")
	fp := filepath.Join("templates", tmplName)

	info, err := os.Stat(fp)
	if err != nil || info.IsDir() {
		renderErrorTemplate(w, http.StatusNotFound, "errors/404.html")
		return
	}

	tmpl, err := template.ParseFiles(lp, fp)
	if err != nil {
		log.Println(err)
		renderErrorTemplate(w, http.StatusInternalServerError, "errors/500.html")
		return
	}

	if err := tmpl.ExecuteTemplate(w, "layout", data); err != nil {
		log.Println(err)
		renderErrorTemplate(w, http.StatusInternalServerError, "errors/500.html")
	}
}

func ServeTemplate(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	if r.URL.Path == "/index.html" {
		http.Redirect(w, r, "/", http.StatusFound)
		return
	}

	lp := filepath.Join("templates", "layout.html")
	var fp string

	if r.URL.Path == "/" {
		fp = filepath.Join("templates", "index.html")
	} else {
		fp = filepath.Join("templates", filepath.Clean(r.URL.Path))
	}

	info, err := os.Stat(fp)
	if err != nil || info.IsDir() {
		renderErrorTemplate(w, http.StatusNotFound, "errors/404.html")
		return
	}

	tmpl, err := template.ParseFiles(lp, fp)
	if err != nil {
		renderErrorTemplate(w, http.StatusInternalServerError, "errors/500.html")
		return
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	if err := tmpl.ExecuteTemplate(w, "layout", nil); err != nil {
		renderErrorTemplate(w, http.StatusInternalServerError, "errors/500.html")
	}
}

func HandleAsciiArt(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodPost {
        http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
        return
    }

    var artStyle, userText string

    ct := r.Header.Get("Content-Type")

if strings.HasPrefix(ct, "application/x-www-form-urlencoded") {
    if err := r.ParseForm(); err != nil {
        http.Error(w, "Invalid form data", http.StatusBadRequest)
        return
    }
    artStyle = r.FormValue("artstyle")
    userText = r.FormValue("text")

} else if strings.HasPrefix(ct, "application/json") {
    var body struct {
        ArtStyle string `json:"artstyle"`
        Text     string `json:"text"`
    }
    if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
        http.Error(w, "Invalid JSON", http.StatusBadRequest)
        return
    }
    artStyle = body.ArtStyle
    userText = body.Text

} else {
    http.Error(w, "Unsupported Content-Type", http.StatusBadRequest)
    return
}


    if artStyle == "" || userText == "" {
        http.Error(w, "Missing art style or text", http.StatusBadRequest)
        return
    }

    artStylePath := filepath.Join("ascii-art", "artstyles", artStyle+".txt")
    if _, err := os.Stat(artStylePath); err != nil {
        http.Error(w, "Art style file not found", http.StatusInternalServerError)
        return
    }

    asciiArtResult := ascii_art.AsciiArt(userText, artStylePath)

    data := struct {
        ASCIIArtResult string
    }{
        ASCIIArtResult: asciiArtResult,
    }

    renderTemplateWithData(w, "index.html", data)
}


func renderErrorTemplate(w http.ResponseWriter, statusCode int, templatePath string) {
	w.WriteHeader(statusCode)

	lp := filepath.Join("templates", "layout.html")
	fp := filepath.Join("templates", templatePath)

	info, err := os.Stat(fp)
	if err != nil || info.IsDir() {
		http.Error(w, http.StatusText(statusCode), statusCode)
		return
	}

	tmpl, err := template.ParseFiles(lp, fp)
	if err != nil {
		http.Error(w, http.StatusText(statusCode), statusCode)
		return
	}

	if err := tmpl.ExecuteTemplate(w, "layout", nil); err != nil {
		http.Error(w, http.StatusText(statusCode), statusCode)
	}
}
