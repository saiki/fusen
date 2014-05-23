package main

import (
	"encoding/json"
	"github.com/saiki/petapeta/wall"
	"net/http"
)

func init() {
	http.Handle("/", http.FileServer(http.Dir("static")))
	http.HandleFunc("/all", all)
	http.HandleFunc("/create", create)
	http.HandleFunc("/update", update)
	http.HandleFunc("/remove", remove)
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

func create(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	fusen := new(wall.Fusen)
	err := decoder.Decode(fusen)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	index, err := whiteboard.Add(fusen)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	p := new(param)
	p.index = index
	p.fusen = fusen
	js, err := json.Marshal(p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

type param struct {
	index string
	fusen *wall.Fusen
}

func update(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	p := new(param)
	err := decoder.Decode(p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	err = whiteboard.Modify(p.index, p.fusen)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	js, err := json.Marshal(p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

func remove(w http.ResponseWriter, r *http.Request) {
	index := r.PostForm.Get("index")
	err := whiteboard.Delete(index)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.WriteHeader(200)
}
