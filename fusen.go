package main

import (
	"fmt"
	"os"
	"github.com/saiki/fusen/bookshelf"
)


func main() {
	var note = new(bookshelf.Note).Init()
	page := new(bookshelf.Page).Init()
	pageIndex := note.AddPage(page)
	fusen := bookshelf.NewFusen(0, 0, "#000000", "テスト")
	note.AddFusen(fusen, pageIndex)
	fmt.Printf("%q\n", os.Args)
}
