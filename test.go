package main

import (
	"fmt"
	"gutbot/engine"
)

func main() {
	var position engine.Position
	var move engine.Move

	position.FromFEN("rnbqkbnr/ppp1ppp1/8/3pP2p/8/8/PPPP1PPP/RNBQKBNR w KQkq d6 0 3")
	fmt.Print(position.String())

	move.FromUCI("e5d6")
	position.Make(move)
	fmt.Print((position.String()))
}
