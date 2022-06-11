package engine

import (
	"fmt"
	"strconv"
	"strings"
)

const (
	// SIDE
	white  = 0
	black  = 1
	noSide = 2

	// PIECE
	pawn   = 0
	knight = 1
	bishop = 2
	rook   = 3
	queen  = 4
	king   = 5
	noKind = 6

	// MOVE TYPE
	normal  = 0
	enpass  = 1
	castle  = 2
	promote = 3

	noPiece = "-"
)

type Piece struct {
	kind int
	side int
}

var NO_PIECE = Piece{noKind, noSide}

type Castling struct {
	whiteQueen Piece
	whiteKing  Piece
	blackQueen Piece
	blackKing  Piece
}

type Position struct {
	board [64]Piece

	turn      int
	castling  Castling
	enpassant int
	halfmoves int
	fullmove  int
}

var PieceToChar map[Piece]string = map[Piece]string{
	{pawn, white}:   "P",
	{pawn, black}:   "p",
	{knight, white}: "N",
	{knight, black}: "n",
	{bishop, white}: "B",
	{bishop, black}: "b",
	{rook, white}:   "R",
	{rook, black}:   "r",
	{queen, white}:  "Q",
	{queen, black}:  "q",
	{king, white}:   "K",
	{king, black}:   "k",
}

var CharToPiece map[byte]Piece = map[byte]Piece{
	'P': {pawn, white},
	'p': {pawn, black},
	'N': {knight, white},
	'n': {knight, black},
	'B': {bishop, white},
	'b': {bishop, black},
	'R': {rook, white},
	'r': {rook, black},
	'Q': {queen, white},
	'q': {queen, black},
	'K': {king, white},
	'k': {king, black},
}

var IndexToSquare [64]string = [64]string{
	"a8", "b8", "c8", "d8", "e8", "f8", "g8", "h8",
	"a7", "b7", "c7", "d7", "e7", "f7", "g7", "h7",
	"a6", "b6", "c6", "d6", "e6", "f6", "g6", "h6",
	"a5", "b5", "c5", "d5", "e5", "f5", "g5", "h5",
	"a4", "b4", "c4", "d4", "e4", "f4", "g4", "h4",
	"a3", "b3", "c3", "d3", "e3", "f3", "g3", "h3",
	"a2", "b2", "c2", "d2", "e2", "f2", "g2", "h2",
	"a1", "b1", "c1", "d1", "e1", "f1", "g1", "h1",
}

func SquareToIndex(square string) int {
	for i, v := range IndexToSquare {
		if v == square {
			return i
		}
	}
	return -1
}

func (position *Position) decideMoveKind(move Move, moving Piece) int {
	// enpass
	if move.to == position.enpassant {
		return enpass
	}
	// castling
	if moving.kind == king {
		if moving.side == white && IndexToSquare[move.to] == "g1" && position.castling.whiteKing != NO_PIECE {
			return castle
		}
		if moving.side == white && IndexToSquare[move.to] == "c1" && position.castling.whiteQueen != NO_PIECE {
			return castle
		}
		if moving.side == black && IndexToSquare[move.to] == "g8" && position.castling.blackKing != NO_PIECE {
			return castle
		}
		if moving.side == black && IndexToSquare[move.to] == "c8" && position.castling.blackQueen != NO_PIECE {
			return castle
		}
	}
	return normal //TODO IMPLEMENT THE OTHERS
}

func (position *Position) Make(move Move) {
	to := move.to
	from := move.from
	moving := position.board[from]
	moveKind := position.decideMoveKind(move, moving)

	switch moveKind {
	case normal:
		position.remove(from)
		position.remove(to)
		position.place(moving, to)
	case enpass:
		position.place(moving, to)
		position.remove(from)
		if moving.side == black {
			position.remove(to - 8)
		} else {
			position.remove(to + 8)
		}
	case castle:
		if moving.side == white && IndexToSquare[move.to] == "g1" {
			position.remove(from)
			rook := position.getPieceAt(SquareToIndex("h1"))
			position.remove(SquareToIndex("h1"))
			position.place(rook, SquareToIndex("f1"))
			position.place(moving, to)
		} else if moving.side == white && IndexToSquare[move.to] == "c1" {
			position.remove(from)
			rook := position.getPieceAt(SquareToIndex("a1"))
			position.remove(SquareToIndex("a1"))
			position.place(rook, SquareToIndex("d1"))
			position.place(moving, to)
		} else if moving.side == black && IndexToSquare[move.to] == "g8" {
			position.remove(from)
			rook := position.getPieceAt(SquareToIndex("h8"))
			position.remove(SquareToIndex("h8"))
			position.place(rook, SquareToIndex("f8"))
			position.place(moving, to)
		} else if moving.side == black && IndexToSquare[move.to] == "c8" {
			position.remove(from)
			rook := position.getPieceAt(SquareToIndex("a8"))
			position.remove(SquareToIndex("a8"))
			position.place(rook, SquareToIndex("d8"))
			position.place(moving, to)
		}
	case promote:
		// TODO
	}

	if position.turn == white {
		position.turn = black
	} else if position.turn == black {
		position.turn = white
	}
}

func (position *Position) place(piece Piece, square int) {
	position.board[square] = piece
}

func (position *Position) remove(square int) {
	position.board[square].kind = noKind
	position.board[square].side = noSide
}

func (position *Position) getPieceAt(square int) Piece {
	return position.board[square]
}

func coordToIndex(row int, col int) int {
	return 8*row + col
}

func (position *Position) String() (result string) {
	result += "\n"
	for row := 0; row < 8; row++ {
		for col := 0; col < 8; col++ {
			if position.board[coordToIndex(row, col)].kind != noKind {
				result += PieceToChar[position.board[coordToIndex(row, col)]]
			} else {
				result += noPiece
			}
		}
		result += "\n"
	}
	return result
}

func (position *Position) ToFEN() string {
	fen := "\n"
	// pieces
	for row := 0; row < 8; row++ {
		empties := 0
		for col := 0; col < 8; col++ {
			if position.board[coordToIndex(row, col)].kind == noKind {
				empties++
				continue
			} else {
				fen += PieceToChar[position.board[coordToIndex(row, col)]]
			}
		}
		if empties > 0 {
			fen += fmt.Sprint(empties)
			empties = 0
		}
		fen += "/"
	}
	fen += " "
	// turn
	if position.turn == white {
		fen += "b"
	} else {
		fen += "w"
	}
	fen += " "
	// castling
	fen += PieceToChar[position.castling.whiteKing]
	fen += PieceToChar[position.castling.whiteQueen]
	fen += PieceToChar[position.castling.blackKing]
	fen += PieceToChar[position.castling.blackQueen]
	fen += " "
	// en passant
	if position.enpassant != -1 {
		fen += IndexToSquare[position.enpassant]
	} else {
		fen += "-"
	}

	fen += " "

	// halfmove clock
	fen += strconv.Itoa(position.halfmoves)
	fen += " "
	// fullmove clock
	fen += strconv.Itoa(position.fullmove)
	fen += " "
	return fen
}

func byteDigitToInt(b byte) int {
	str := string(b)
	number, err := strconv.Atoi(str)
	if err == nil {
		return number
	}
	return 0
}

func (position *Position) FromFEN(fen string) {
	sections := strings.Split(fen, " ")
	board := sections[0]
	turn := sections[1]
	castling := sections[2]
	enpassant := sections[3]
	halfmove := sections[4]
	fullmove := sections[5]

	pos := 0
	// board
	for char := range board {
		piece := board[char]
		switch piece {
		case 'P', 'p', 'N', 'n', 'B', 'b', 'R', 'r', 'Q', 'q', 'K', 'k':
			{
				position.place(CharToPiece[piece], pos)
				pos++
				break
			}
		case '/':
			{
				break
			}
		case '1', '2', '3', '4', '5', '6', '7', '8':
			{
				empties := byteDigitToInt(piece)
				for i := 0; i < empties; i++ {
					position.remove(pos)
					pos++
				}
			}
		}
	}
	// turn
	if turn == "w" {
		position.turn = white
	} else {
		position.turn = black
	}
	// enpassant
	position.enpassant = SquareToIndex(enpassant)
	// castling
	position.castling.whiteKing = CharToPiece[castling[0]]
	position.castling.whiteQueen = CharToPiece[castling[1]]
	position.castling.blackKing = CharToPiece[castling[2]]
	position.castling.blackQueen = CharToPiece[castling[3]]
	// halfmove clock
	halfmoves, err := strconv.Atoi(halfmove)
	if err != nil {
		position.halfmoves = halfmoves
	}
	// hi matthew gurski :)
	// fullmove clock
	fullmoves, err := strconv.Atoi(fullmove)
	if err != nil {
		position.fullmove = fullmoves
	}
}
