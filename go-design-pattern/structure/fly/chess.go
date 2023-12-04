package fly

type ChessUnit struct {
	id    int
	color string
	name  string
}

var chessUnits = map[int]*ChessUnit{
	1: {id: 1, color: "红", name: "车"},
	2: {id: 2, color: "黑", name: "车"},
}

type ChessPiece struct {
	Unit *ChessUnit
	X    int
	Y    int
}

type ChessBoard struct {
	ChessPieces map[int]*ChessPiece
}

func NewChessBoard() *ChessBoard {
	board := &ChessBoard{ChessPieces: map[int]*ChessPiece{}}
	for id := range chessUnits {
		board.ChessPieces[id] = &ChessPiece{
			Unit: chessUnits[id],
			X:    0,
			Y:    0,
		}
	}
	return board
}

func (c *ChessBoard) Move(id, x, y int) {
	c.ChessPieces[id].X = x
	c.ChessPieces[id].Y = y
}
