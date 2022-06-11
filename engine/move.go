package engine

type Move struct {
	to        int
	from      int
	promoteTo Piece
}

func (move *Move) FromUCI(uci string) {
	from := uci[0:2]
	to := uci[2:4]
	move.to = SquareToIndex(to)
	move.from = SquareToIndex(from)

	if len(uci) == 5 {
		promotePiece := []byte(uci[4:])
		move.promoteTo = CharToPiece[promotePiece[0]]
	} else {
		move.promoteTo = NO_PIECE
	}
}
