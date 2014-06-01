package wall

import (
	"bytes"
	"encoding/gob"
	"errors"
	"log"
	"os"
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
	Collection []*Fusen
}

func (self *Wall) Init() *Wall {
	self.Collection = make([]*Fusen, 0)
	return self
}

func (self *Wall) Add(fusen *Fusen) (int, error) {
	_ = append(self.Collection, fusen)
	return len(self.Collection), nil
}

func (self *Wall) Modify(index int, fusen *Fusen) error {
	length := len(self.Collection)
	if length - index < 0 {
		return errors.New("index out of bounds.")
	}
	self.Collection[index] = fusen
	return nil
}

func (self *Wall) Delete(index int) error {
	length := len(self.Collection)
	if length - index < 0 {
		return errors.New("index out of bounds.")
	}
	self.Collection = append(self.Collection[:index], self.Collection[index+1:]...)
	c := make([]*Fusen, len(self.Collection) - 1)
	copy(c, self.Collection)
	self.Collection = c
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
