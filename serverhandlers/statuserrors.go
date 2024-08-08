package serverhandlers

import (
	"html/template"
	"log"
	"net/http"
)

func StatusError(w http.ResponseWriter, r *http.Request, statusCode int, templatePath, message string) {
	tmpl, err := template.ParseFiles(templatePath)
	if err != nil {
		http.Error(w, message, statusCode)
		log.Printf("Error parsing %s: %v", templatePath, err)
		return
	}
	w.WriteHeader(statusCode)
	err = tmpl.Execute(w, nil)
	if err != nil {
		http.Error(w, message, statusCode)
		log.Printf("Error executing template %s: %v", templatePath, err)
	}
}

func Serve404File(w http.ResponseWriter, r *http.Request) {
	StatusError(w, r, http.StatusNotFound, "templates/404file.html", "File Not Found")
}

func Serve404Page(w http.ResponseWriter, r *http.Request) {
	StatusError(w, r, http.StatusNotFound, "templates/404page.html", "Page Not Found")
}

func Serve400(w http.ResponseWriter, r *http.Request) {
	StatusError(w, r, http.StatusBadRequest, "templates/400page.html", "Bad Request")
}

func Serve405(w http.ResponseWriter, r *http.Request) {
	StatusError(w, r, http.StatusMethodNotAllowed, "templates/405page.html", "Method Not Allowed")
}

func Serve500(w http.ResponseWriter, r *http.Request) {
	StatusError(w, r, http.StatusInternalServerError, "templates/500page.html", "Something Unexpected Happened. Try Again Later")
}
