package engine

const (
	normal  = 0
	enpass  = 1
	castle  = 2
	promote = 3
)

type Move struct {
	to        int
	from      int
	capturing int
	player    int
	moveKind  int
}
