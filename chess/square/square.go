package square

import (
	"chess/chess/axis"
	"chess/chess/piece"
)

type Square interface {
	SetPiece(piece piece.Piece)
	Piece() piece.Piece
	UnsetPiece()
}

type square struct {
	x     axis.X
	y     axis.Y
	piece piece.Piece
}

func New(x axis.X, y axis.Y) Square {
	return &square{
		x: x,
		y: y,
	}
}

func (s *square) SetPiece(piece piece.Piece) {
	s.piece = piece
}

func (s *square) Piece() piece.Piece {
	return s.piece
}

func (s *square) UnsetPiece() {
	s.piece = nil
}
