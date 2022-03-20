package main

import (
	"embed"
	"io/fs"
	"log"
	"net/http"
)

//go:embed templates/*
var templates embed.FS

func main() {
	//setupLogFile()
	log.Println("Info :", "Starting ScoreBoardApp web app")
	staticFiles, err := fs.Sub(fs.FS(templates), "templates/assets")
	if err != nil {
		log.Println("unable to load static files", err)
	}
	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.FS(staticFiles))))

	http.HandleFunc("/", handleScore)
	http.HandleFunc("/scoreboardview", handleScoreBoard)
	http.HandleFunc("/scoreview", handleScore)
	http.HandleFunc("/savescore", handleSaveScore)
	log.Fatal(http.ListenAndServe(":"+"8080", nil))
}
