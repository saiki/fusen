package main

import (
	"encoding/json"
//	"github.com/saiki/petapeta/wall"
	"./model"
	"net/http"
	"strconv"
)

func init() {
	http.Handle("/", http.FileServer(http.Dir("static")))
	http.HandleFunc("/all", all)
	http.HandleFunc("/create", create)
	http.HandleFunc("/update", update)
	http.HandleFunc("/remove", remove)
}

func all(w http.ResponseWriter, r *http.Request) {
	js, err := json.Marshal(collection)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

func create(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	fusen := new(model.Fusen)
	err := decoder.Decode(fusen)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	_, err = collection.Add(fusen)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	all(w, r)
}

type param struct {
	index int
	fusen *model.Fusen
}

func update(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	p := new(param)
	err := decoder.Decode(p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	err = collection.Modify(p.index, p.fusen)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	all(w, r)
}

func remove(w http.ResponseWriter, r *http.Request) {
	index, err := strconv.Atoi(r.PostForm.Get("index"))
	if ( err != nil ) {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	err = collection.Delete(index)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	all(w, r)
}
