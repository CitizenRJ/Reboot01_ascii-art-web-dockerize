package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	// "os"
	"asciiartweb/internal/asciiart"
	"asciiartweb/internal/asciiartfs"
)

type Fonts struct {
	Art    string
	Hidden string
}

const port = ":8080"

func printHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		parsedTemplate, _ := template.ParseFiles("../../static/405.html")
		parsedTemplate.Execute(w, nil)
		return
	}

	if r.URL.Path != "/" {
		w.WriteHeader(http.StatusNotFound)
		parsedTemplate, _ := template.ParseFiles("../../static/404.html")
		parsedTemplate.Execute(w, nil)
		return
	}

	parsedTemplate, err := template.ParseFiles("../../static/index.html")
	if err != nil {
		log.Println("Error executing template :", err)
		return
	}

	err = parsedTemplate.Execute(w, nil)
	if err != nil {
		log.Println("Error executing template :", err)
		return
	}
}

func formHandler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return
	}
	if r.Method != "POST" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		parsedTemplate, _ := template.ParseFiles("../../static/405.html")
		parsedTemplate.Execute(w, nil)
		return
	} else {
		name := r.FormValue("name")
		errName := asciiartfs.IsValid(name)
		if errName != nil {
			w.WriteHeader(http.StatusBadRequest)
			parsedTemplate, _ := template.ParseFiles("../../static/400.html")
			err := parsedTemplate.Execute(w, nil)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				parsedTemplate, _ := template.ParseFiles("../../static/500.html")
				err = parsedTemplate.Execute(w, nil)
				if err != nil {
					http.Error(w, "Internal Server Error", http.StatusInternalServerError)
				}
			}
			return
		}

		banner := r.FormValue("banner")
		if banner != "shadow" && banner != "standard" && banner != "thinkertoy" {
			w.WriteHeader(http.StatusNotFound)
			parsedTemplate, _ := template.ParseFiles("../../static/404.html")
			parsedTemplate.Execute(w, nil)
			return
		}
		fmt.Println(banner)
		art, err := asciiart.AsciiArt(banner, name)
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			parsedTemplate, _ := template.ParseFiles("../../static/404.html")
			parsedTemplate.Execute(w, nil)
			return
		}
		fonts := Fonts{Art: art, Hidden: "false"}
		parsedTemplate, err := template.ParseFiles("../../static/index.html")
		if err != nil {
			log.Println("Error executing template :", err)
			return
		}

		err = parsedTemplate.Execute(w, fonts)
		if err != nil {
			log.Println("Error executing template :", err)
			return
		}
	}
}

func main() {
	http.HandleFunc("/", printHandler)
	http.HandleFunc("/ascii-art", formHandler)
	http.HandleFunc("/w.css", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "../../static/w.css")
	})

	fmt.Printf("Starting server at http://localhost" + port + "/\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
