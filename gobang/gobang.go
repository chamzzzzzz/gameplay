package gobang

import (
	"fmt"
)

const (
	Empty   = 0
	Black   = 1
	White   = 2
	Draw    = 3
	Invalid = -1
)

type Position struct {
	Row, Col int
}

type Board struct {
	size    int
	chesses [][]int
}

func NewBoard(size int) *Board {
	chesses := make([][]int, size)
	for i := range chesses {
		chesses[i] = make([]int, size)
	}
	return &Board{
		size:    size,
		chesses: chesses,
	}
}

func (b *Board) Reset() {
	for i := range b.chesses {
		for j := range b.chesses[i] {
			b.chesses[i][j] = Empty
		}
	}
}

func (b *Board) Size() int {
	return b.size
}

func (b *Board) Get(row, col int) int {
	if row < 0 || row >= b.size || col < 0 || col >= b.size {
		return Invalid
	}
	return b.chesses[row][col]
}

func (b *Board) Set(row, col, chess int) {
	if row < 0 || row >= b.size || col < 0 || col >= b.size {
		return
	}
	b.chesses[row][col] = chess
}

type Result struct {
	Winner    int
	Positions []Position
}

func (r *Result) Reset() {
	r.Winner = Empty
	r.Positions = nil
}

func (r *Result) AddPosition(row, col int) {
	pos := Position{Row: row, Col: col}
	for _, p := range r.Positions {
		if p == pos {
			return
		}
	}
	r.Positions = append(r.Positions, pos)
}

type Gameplay struct {
	board *Board
	turn  int
}

func NewGameplay(board *Board, turn int) *Gameplay {
	if board == nil {
		board = NewBoard(15)
	}
	if turn != Black && turn != White {
		turn = Black
	}
	return &Gameplay{
		board: board,
		turn:  Black,
	}
}

func (g *Gameplay) Reset() {
	g.board.Reset()
	g.turn = Black
}

func (g *Gameplay) Turn() int {
	return g.turn
}

func (g *Gameplay) Move(row, col int) error {
	r := g.Judge()
	if r.Winner != Empty {
		return fmt.Errorf("invalid move: game over")
	}
	size := g.board.Size()
	if row < 0 || row > size {
		return fmt.Errorf("invalid move: row index must be between 0 and %d", size-1)
	}
	if col < 0 || col > size {
		return fmt.Errorf("invalid move: column index must be between 0 and %d", size-1)
	}
	if g.board.Get(row, col) != Empty {
		return fmt.Errorf("invalid move: not empty")
	}
	g.board.Set(row, col, g.turn)
	if g.turn == Black {
		g.turn = White
	} else {
		g.turn = Black
	}
	return nil
}

func (g *Gameplay) Judge() (result Result) {
	size := g.board.Size()
	count := 0
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			chess := g.board.Get(i, j)
			if chess == Empty {
				continue
			}
			count++

			var vertical, horizontal, rightUp, rightDown int
			for k := 1; k < 5; k++ {
				if i+k < size {
					if chess == g.board.Get(i+k, j) {
						horizontal++
					}
				}
				if j+k < size {
					if chess == g.board.Get(i, j+k) {
						vertical++
					}
				}
				if i+k < size && j-k >= 0 {
					if chess == g.board.Get(i+k, j-k) {
						rightUp++
					}
				}
				if i+k < size && j+k < size {
					if chess == g.board.Get(i+k, j+k) {
						rightDown++
					}
				}
			}

			if horizontal == 4 || vertical == 4 || rightUp == 4 || rightDown == 4 {
				if result.Winner == Empty {
					result.Winner = chess
				} else if result.Winner != chess {
					result.Reset()
					result.Winner = Invalid
					return
				}
				result.AddPosition(i, j)
				if horizontal == 4 {
					for k := 1; k < 5; k++ {
						result.AddPosition(i+k, j)
					}
				}
				if vertical == 4 {
					for k := 1; k < 5; k++ {
						result.AddPosition(i, j+k)
					}
				}
				if rightUp == 4 {
					for k := 1; k < 5; k++ {
						result.AddPosition(i+k, j-k)
					}
				}
				if rightDown == 4 {
					for k := 1; k < 5; k++ {
						result.AddPosition(i+k, j+k)
					}
				}
			}
		}
	}
	if result.Winner == Empty && count == size*size {
		result.Winner = Draw
	}
	return
}

func (g *Gameplay) Board() *Board {
	return g.board
}
