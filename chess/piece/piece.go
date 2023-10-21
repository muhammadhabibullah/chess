package piece

import "chess/chess/color"

type Piece interface {
	Color() color.Color
	Type() Type
	Captured()
	IsCaptured() bool
}

type piece struct {
	color      color.Color
	theType    Type
	isCaptured bool
}

func New(color color.Color, theType Type) Piece {
	return &piece{
		color:   color,
		theType: theType,
	}
}

func (p *piece) Color() color.Color {
	return p.color
}

func (p *piece) Type() Type {
	return p.theType
}

func (p *piece) Captured() {
	p.isCaptured = true
}

func (p *piece) IsCaptured() bool {
	return p.isCaptured
}
