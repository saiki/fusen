package main

import (
	"fmt"
	"os"
	"github.com/saiki/fusen/whiteboard"
)


func main() {
	var board = new(whiteboard.Whiteboard).Init()
	fusen := whiteboard.NewFusen(0, 0, "#000000", "テスト")
	board.Add(fusen)
	fmt.Printf("%v\n", os.Args)
	fmt.Printf("%v\n", board)
	whiteboard.Export(board)
}
