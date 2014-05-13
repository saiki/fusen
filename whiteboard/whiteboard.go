package whiteboard

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"log"
)

type Fusen struct {
	Left  int
	Top   int
	Color string
	Body  string
}

func NewFusen(left, top int, color, body string) *Fusen {
	return &Fusen{Left: left, Top: top, Color: color, Body:body}
}

type Whiteboard struct {
	Collection map[int]*Fusen
}

func (self *Whiteboard) Init() *Whiteboard {
	self.Collection = make(map[int]*Fusen)
	return self
}

func (self *Whiteboard) Add(fusen *Fusen) int {
	next := len(self.Collection) + 1
	self.Collection[next] = fusen
	return next
}

func Export(whiteboard *Whiteboard) error {
	var network bytes.Buffer
	enc := gob.NewEncoder(&network)
	err := enc.Encode(whiteboard)
	if  err != nil {
		log.Fatalln("decode: ", err)
	}
	fmt.Printf("network: %v\n", network)
	return nil
}

func Import(path string) (*Whiteboard, error) {

	return nil, nil
}
