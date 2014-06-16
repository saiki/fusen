package model

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

func NewFusen(left, top, width, height int, color, body string) Fusen {
	return Fusen{Left: left, Top: top, Width: width, Height: height, Color: color, Body:body}
}

type Collection []Fusen


func (self *Collection) Add(fusen *Fusen) (int, error) {
	_ = append(*self, *fusen)
	return len(*self), nil
}

func (self *Collection) Modify(index int, fusen *Fusen) error {
	length := len(*self)
	if length - index < 0 {
		return errors.New("index out of bounds.")
	}
	(*self)[index] = *fusen
	return nil
}

func (self *Collection) Delete(index int) error {
	length := len(*self)
	if length - index < 0 {
		return errors.New("index out of bounds.")
	}
	*self = append((*self)[:index], (*self)[index+1:]...)
	c := make(Collection, len(*self) - 1)
	copy(c, *self)
	self = &c
	return nil
}

func (self *Collection) Export(path string) error {
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

func (self *Collection) Import (path string) error {
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
