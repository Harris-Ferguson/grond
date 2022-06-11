package main

import (
	"fmt"
	"gutbot/engine"
)

func main() {
	var position engine.Position
	var move engine.Move

	fmt.Print("White Kingside")
	position.FromFEN("r3k2r/pppq1ppp/2npbn2/2b1p3/2B1P3/2NPBN2/PPPQ1PPP/R3K2R w KQkq - 6 8")
	fmt.Print(position.String())
	move.FromUCI("e1g1")
	position.Make(move)
	fmt.Print((position.String()))

	fmt.Print("White Queenside")
	position.FromFEN("r3k2r/pppq1ppp/2npbn2/2b1p3/2B1P3/2NPBN2/PPPQ1PPP/R3K2R w KQkq - 6 8")
	fmt.Print(position.String())
	move.FromUCI("e1c1")
	position.Make(move)
	fmt.Print((position.String()))

	fmt.Print("Black Kingside")
	position.FromFEN("r3k2r/pppq1ppp/2npbn2/2b1p3/2B1P3/2NPBN2/PPPQ1PPP/R3K2R w KQkq - 6 8")
	fmt.Print(position.String())
	move.FromUCI("e8g8")
	position.Make(move)
	fmt.Print((position.String()))

	fmt.Print("Black Queenside")
	position.FromFEN("r3k2r/pppq1ppp/2npbn2/2b1p3/2B1P3/2NPBN2/PPPQ1PPP/R3K2R w KQkq - 6 8")
	fmt.Print(position.String())
	move.FromUCI("e8c8")
	position.Make(move)
	fmt.Print((position.String()))

	fmt.Print("Promoting Piece")
	position.FromFEN("8/1K5k/6pP/6P1/8/8/4p3/8 b - - 1 5")
	fmt.Print(position.String())
	move.FromUCI("e2e1q")
	position.Make(move)
	fmt.Print(position.String())
}
