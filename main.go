package main

import (
	"fmt"
	"os"
)

type Fusen struct {
	page  int
	left  int
	top   int
	color string
	body  string
}

type Storage struct {
	data map[int] Fusen
}

func (self *Storage) Init() *Storage {
	self.data = make(map[int] Fusen)
	return self
}

func (self *Storage) Add(fusen Fusen) {
	self.data[fusen.page] = fusen
}

func main() {
	var storage = new(Storage).Init()
	var fusen   = new(Fusen)
	fusen.page  = 1
	fusen.left  = 0
	fusen.top   = 0
	fusen.color = "#000000"
	fusen.body  = "テスト"
	storage.Add(*fusen)
	fmt.Printf("%q\n", os.Args)
}
