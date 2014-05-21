package main

import (
	"encoding/json"
	"log"
	"github.com/saiki/fusen/wall"
	"net/http"
)

var whiteboard *wall.Wall

const exported = "~/.exported"

func init() {
	whiteboard = new(wall.Wall).Init()
	whiteboard.Import(exported)
	http.Handle("/", http.FileServer(http.Dir("static")))
	http.HandleFunc("/all", all)
}

func main() {
	log.Println("start fusen...")
	defer func() {
		whiteboard.Export(exported)
		log.Println("end fusen...")
	}()
	http.ListenAndServe(":8080", nil)
}

func all(w http.ResponseWriter, r *http.Request) {
	js, err := json.Marshal(whiteboard)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}
