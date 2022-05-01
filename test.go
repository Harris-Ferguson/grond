package main

import (
	"fmt"
	"gutbot/engine"
)

func main() {
	var position engine.Position
	position.FromFEN("rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1")
	fmt.Print(position.String())
	fmt.Print(position.ToFEN())
}
