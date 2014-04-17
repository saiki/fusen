package main

import (
	"fmt"
	"os"
)

type Fusen struct {
	left  int
	top   int
	color string
	body  string
}

type Page struct {
	fusen map[int] *Fusen
}

func (self *Page) Init() *Page {
	self.fusen = make(map[int] *Fusen)
	return self
}

func (self *Page) Add(fusen *Fusen) int {
	next := len(self.fusen) + 1
	self.fusen[next] = fusen
	return next
}

type Storage struct {
	pages map[int] *Page
}

func (self *Storage) Init() *Storage {
	self.pages = make(map[int] *Page)
	return self
}

func (self *Storage) AddPage(page *Page) int {
	next := len(self.pages) + 1
	self.pages[next] = page
	return next
}

func (self *Storage) AddFusen(fusen *Fusen, page int) int {
	index := self.pages[page].Add(fusen)
	return index
}

func main() {
	var storage = new(Storage).Init()
	page := new(Page).Init()
	pageIndex := storage.AddPage(page)
	fusen := new(Fusen)
	fusen.left  = 0
	fusen.top   = 0
	fusen.color = "#000000"
	fusen.body  = "テスト"
	storage.AddFusen(fusen, pageIndex)
	fmt.Printf("%q\n", os.Args)
}
