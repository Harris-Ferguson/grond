package engine

type Move struct {
	to   int
	from int
}

func (move *Move) FromUCI(uci string) {
	from := uci[0:2]
	to := uci[2:4]
	move.to = SquareToIndex(to)
	move.from = SquareToIndex(from)
}
