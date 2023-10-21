package board

import (
	"chess/chess/axis"
	"chess/chess/color"
	"chess/chess/piece"
	"chess/chess/square"
	"fmt"
)

type Board struct {
	squares map[axis.X]map[axis.Y]square.Square
}

func New() Board {
	board := Board{
		squares: make(map[axis.X]map[axis.Y]square.Square),
	}
	xAxes := []axis.X{
		axis.XA,
		axis.XB,
		axis.XC,
		axis.XD,
		axis.XE,
		axis.XF,
		axis.XG,
		axis.XH,
	}
	yAxes := []axis.Y{
		axis.Y1,
		axis.Y2,
		axis.Y3,
		axis.Y4,
		axis.Y5,
		axis.Y6,
		axis.Y7,
		axis.Y8,
	}
	for _, xAxis := range xAxes {
		board.squares[xAxis] = make(map[axis.Y]square.Square)
		for _, yAxis := range yAxes {
			board.squares[xAxis][yAxis] = square.New(xAxis, yAxis)
			if initPiece := InitPiece(xAxis, yAxis); initPiece != nil {
				board.squares[xAxis][yAxis].SetPiece(InitPiece(xAxis, yAxis))
			}
		}
	}
	return board
}

func InitPiece(x axis.X, y axis.Y) piece.Piece {
	pieceColor := color.White
	if y == axis.Y7 || y == axis.Y8 {
		pieceColor = color.Black
	}
	switch y {
	case axis.Y2, axis.Y7:
		return piece.New(pieceColor, piece.Pawn)
	case axis.Y1, axis.Y8:
		switch x {
		case axis.XA, axis.XH:
			return piece.New(pieceColor, piece.Rook)
		case axis.XB, axis.XG:
			return piece.New(pieceColor, piece.Knight)
		case axis.XC, axis.XF:
			return piece.New(pieceColor, piece.Bishop)
		case axis.XD:
			return piece.New(pieceColor, piece.Queen)
		case axis.XE:
			return piece.New(pieceColor, piece.King)
		default:
			return nil
		}
	default:
		return nil
	}
}

func (b Board) View(as color.Color) string {
	var boardView string
	var x, y = axis.XA, axis.Y8
	if as == color.Black {
		x, y = axis.XH, axis.Y1
	}
	for y >= axis.Y1 && y <= axis.Y8 {
		var initX = x
		for x >= axis.XA && x <= axis.XH {
			var squarePieceView = " "
			squarePiece := b.squares[x][y].Piece()
			if squarePiece != nil {
				squarePieceView = fmt.Sprintf("%s", squarePiece.Type())
			}
			boardView = fmt.Sprintf("%s[%s]", boardView, squarePieceView)
			if as == color.Black {
				x--
			} else {
				x++
			}
		}
		boardView = fmt.Sprintf("%s%d\n", boardView, y)
		x = initX
		if as == color.White {
			y--
		} else {
			y++
		}
	}

	return boardView
}
