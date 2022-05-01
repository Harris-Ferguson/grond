package engine

import (
	"fmt"
	"strconv"
	"strings"
)

const (
	white  = 0
	black  = 1
	noSide = 2

	pawn   = 0
	knight = 1
	bishop = 2
	rook   = 3
	queen  = 4
	king   = 5
	noKind = 6

	noPiece = "-"
)

type Piece struct {
	kind int
	side int
}

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

var IndexToSquare [64]string = [64]string{"A1", "A2", "A3", "A4", "A5", "A6", "A7", "A8",
	"B1", "B2", "B3", "B4", "B5", "B6", "B7", "B8",
	"C1", "C2", "C3", "C4", "C5", "C6", "C7", "C8",
	"D1", "D2", "D3", "D4", "D5", "D6", "D7", "D8",
	"E1", "E2", "E3", "E4", "E5", "E6", "E7", "E8",
	"F1", "F2", "F3", "F4", "F5", "F6", "F7", "F8",
	"G1", "G2", "G3", "G4", "G5", "G6", "G7", "G8",
	"H1", "H2", "H3", "H4", "H5", "H6", "H7", "H8",
}

func squareToIndex(square string) int {
	for i, v := range IndexToSquare {
		if v == square {
			return i
		}
	}
	return -1
}

func (position *Position) make(move Move) {
	to := move.to
	from := move.from
	moveKind := move.moveKind
	moving := position.board[from]
	capturing := move.capturing

	switch moveKind {
	case normal:
		position.remove(from)
		position.remove(to)
		position.place(moving, to)
	case enpass:
		position.remove(from)
		position.remove(capturing)
		position.place(moving, to)
	case castle:
	case promote:
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
	position.enpassant = squareToIndex(enpassant)
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
