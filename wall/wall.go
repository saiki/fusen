package wall

import (
	"bytes"
	"encoding/gob"
	"errors"
	"log"
	"os"
	"strconv"
)

type Fusen struct {
	Left   int
	Top    int
	Width  int
	Height int
	Color  string
	Body   string
}

func NewFusen(left, top, width, height int, color, body string) *Fusen {
	return &Fusen{Left: left, Top: top, Width: width, Height: height, Color: color, Body:body}
}

type Wall struct {
	Collection map[string]*Fusen
}

func (self *Wall) Init() *Wall {
	self.Collection = make(map[string]*Fusen)
	return self
}

func (self *Wall) Add(fusen *Fusen) (string, error) {
	next := len(self.Collection) + 1
	self.Collection[strconv.Itoa(next)] = fusen
	return strconv.Itoa(next), nil
}

func (self *Wall) Modify(index string, fusen *Fusen) error {
	if _, exists := self.Collection[index]; !exists {
		return errors.New("index not found.")
	}
	self.Collection[index] = fusen
	return nil
}

func (self *Wall) Delete(index string) error {
	if _, exists := self.Collection[index]; !exists {
		return errors.New("index not found.")
	}
	delete(self.Collection, index)
	return nil
}

func (self *Wall) Export(path string) error {
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	err := enc.Encode(self)
	if  err != nil {
		return err
	}
	if _, err := os.Stat(path); os.IsExist(err) {
		err := os.Remove(path)
		if err != nil {
			return err
		}
	}
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()
	file.Write(buf.Bytes())
	return nil
}

func (self *Wall) Import (path string) error {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		log.Println("file does not exists.")
		return nil
	}
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	dec := gob.NewDecoder(file)
	err = dec.Decode(self)
	if err != nil {
		return err
	}
	return nil
}
