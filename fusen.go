package main

import (
	"log"
	"github.com/saiki/fusen/whiteboard"
	"net/http"
)

var board *whiteboard.Whiteboard

const exported = "~/.exported"

func init() {
	board = new(whiteboard.Whiteboard).Init()
	board.Import(exported)
}

func main() {
	log.Println("start fusen...")
	defer board.Export(exported)
	http.ListenAndServe(":8080", nil)
}
