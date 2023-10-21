package piece

type Type int

const (
	_ Type = iota
	Pawn
	Knight
	Bishop
	Rook
	Queen
	King
)

func (t Type) String() string {
	switch t {
	case Pawn:
		return "p"
	case Rook:
		return "R"
	case Knight:
		return "N"
	case Bishop:
		return "B"
	case Queen:
		return "Q"
	case King:
		return "K"
	default:
		return ""
	}
}
