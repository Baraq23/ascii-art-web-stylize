package serverhandlers

import (
	"errors"
	"html/template"
	"log"
	"net/http"
	"os"
	"strings"

	"ascii-art-web-stylize/printasciiart"
	"ascii-art-web-stylize/readbanner"
)

type Page struct {
	Input, AsciiArt string
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		Serve404Page(w, r)
		return
	}

	tmpl, err := template.ParseFiles("templates/index.html")
	if err != nil {
		Serve404File(w, r)
		log.Printf("Error parsing index.html: %v\n", err)
		return
	}
	err = tmpl.Execute(w, nil)
	if err != nil {
		Serve500(w, r)
		log.Printf("Error executing template: %v\n", err)
		return
	}
}

func AboutHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		Serve405(w, r)
		return
	}

	tmpl, err := template.ParseFiles("templates/about.html")
	if err != nil {
		Serve404File(w, r)
		log.Printf("Error parsing index.html: %v\n", err)
		return
	}

	err = tmpl.Execute(w, nil)
	if err != nil {
		Serve500(w, r)
		log.Printf("Error executing template: %v\n", err)
		return
	}
}

func AsciiArtHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		Serve405(w, r)
		log.Println("Method to be accessed should be POST")
		return
	}

	// Retrieves the input string and banner file from the request form data
	inputString := r.FormValue("input")
	bannerFile := r.FormValue("banner")

	if inputString == "" {
		Serve400(w, r)
		log.Println("Input cannot be empty")
		return
	}

	if bannerFile == "" {
		bannerFile = "standard"
	}

	bannerFilePath := "banners/" + bannerFile + ".txt"
	bannerFileSlice, err := readbanner.ReadBanner(bannerFilePath)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			Serve404File(w, r)
			log.Println("Banner file not found: " + bannerFilePath)
		} else if errors.Is(err, readbanner.ErrBannerTampered) {
			Serve500(w, r)
			log.Println("Banner file altered: " + bannerFilePath)
		} else {
			Serve500(w, r)
			log.Println("Error with banner file: " + bannerFilePath)
		}
		return
	}

	var output strings.Builder
	// Generates ASCII art using the PrintArt function and writes the output to a strings.Builder
	err = printasciiart.PrintAsciiArt(bannerFileSlice, inputString, &output)
	if err != nil {
		Serve400(w, r)
		return
	}

	data := Page{
		Input:    inputString,
		AsciiArt: output.String(),
	}

	tmpl, err := template.ParseFiles("templates/index.html")
	if err != nil {
		Serve404File(w, r)
		return
	}
	tmpl.Execute(w, data)
}
