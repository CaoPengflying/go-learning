package fly

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewChessBoard(t *testing.T) {
	board := NewChessBoard()
	board.Move(1,1,1)
	board2 := NewChessBoard()

	board2.Move(1,2,2)

	assert.Equal(t, board.ChessPieces[1].Unit,board2.ChessPieces[1].Unit)
	assert.Equal(t,board.ChessPieces[2].Unit,board2.ChessPieces[2].Unit)
}
