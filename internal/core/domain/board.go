package domain

import (
	"math/rand"
	"time"
)

const (
	CELL_BOMB     = "X"
	CELL_EMPTY    = "-"
	CELL_REVEALED = "O"

	CELL_BOMB_HIDED = "-"
)

type Board [][]string

func NewBoard(size uint, bombs uint) Board {
	board := NewEmptyBoard(size)
	board.FillWithBombs(bombs)

	return board
}

func (board Board) FillWithBombs(bombs uint) {
	rows := len(board)
	cols := len(board[0])
	positions := getRandomPositions(rows*cols, bombs)

	var row, col int

	for _, pos := range positions {
		row = pos / cols
		col = pos - row*rows

		board[row][col] = CELL_BOMB
	}
}

func (board Board) HideBombs() Board {
	newBoard := NewEmptyBoard(uint(len(board)))

	for row := range board {
		for col := range board[0] {
			if board[row][col] == CELL_BOMB {
				newBoard[row][col] = CELL_BOMB_HIDED
			} else {
				newBoard[row][col] = board[row][col]
			}
		}
	}

	return newBoard
}

func (board Board) IsValidPosition(row uint, col uint) bool {
	return row < uint(len(board)) && col < uint(len(board[0]))
}

func (board Board) Contains(row uint, col uint, element string) bool {
	return board[row][col] == element
}

func (board Board) Set(row uint, col uint, element string) {
	board[row][col] = element
}

func (board Board) HasEmptyCells() bool {
	for row := range board {
		for col := range board[0] {
			if board[row][col] == CELL_EMPTY {
				return true
			}
		}
	}

	return false
}

func NewEmptyBoard(size uint) Board {
	board := make([][]string, size)
	for row := range board {
		board[row] = make([]string, size)
	}

	for row := range board {
		for col := range board[0] {
			board[row][col] = CELL_EMPTY
		}
	}

	return board
}

// ··· Private Functions ··· //

func getRandomPositions(size int, n uint) []int {
	rand.Seed(time.Now().UnixNano())
	p := rand.Perm(size)

	positions := []int{}

	for _, r := range p[:n] {
		positions = append(positions, r)
	}

	return positions
}
