package main

import (
	"fmt"
	"os"
	"github.com/saiki/fusen/whiteboard"
)


func main() {
	var board = new(whiteboard.Whiteboard).Init()
	board.Import(".\\exported")
	fusen := whiteboard.NewFusen(0, 0, "#000000", "テスト")
	board.Add(fusen)
	fmt.Printf("%v\n", os.Args)
	fmt.Printf("%v\n", board)
	board.Export(".\\exported")
}
