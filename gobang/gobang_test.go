package gobang

import (
	"testing"
)

func TestNewBoard(t *testing.T) {
	b := NewBoard(15)
	if b == nil {
		t.Error("want not nil, got nil")
	}
	if b.Size() != 15 {
		t.Errorf("want 15, got %d", b.Size())
	}
}

func TestBoardSetGet(t *testing.T) {
	b := NewBoard(15)
	size := b.Size()
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			want := Black
			if i%2 == 0 {
				want = White
			}
			b.Set(i, j, want)
			if got := b.Get(i, j); got != want {
				t.Errorf("want %d, got %d", want, got)
			}
		}
	}
}

func TestBoardReset(t *testing.T) {
	b := NewBoard(15)
	size := b.Size()
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			want := Black
			if i%2 == 0 {
				want = White
			}
			b.Set(i, j, want)
		}
	}
	b.Reset()
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			if got := b.Get(i, j); got != Empty {
				t.Errorf("want %d, got %d", Empty, got)
			}
		}
	}
}

func TestBoardSetGetInvalid(t *testing.T) {
	b := NewBoard(15)
	want := Invalid
	b.Set(-1, 0, Black)
	if got := b.Get(-1, 0); got != want {
		t.Errorf("want %d, got %d", want, got)
	}
	b.Set(0, -1, Black)
	if got := b.Get(0, -1); got != want {
		t.Errorf("want %d, got %d", want, got)
	}
	b.Set(15, 0, Black)
	if got := b.Get(15, 0); got != want {
		t.Errorf("want %d, got %d", want, got)
	}
	b.Set(0, 15, Black)
	if got := b.Get(0, 15); got != want {
		t.Errorf("want %d, got %d", want, got)
	}
}

func TestNewGameplay(t *testing.T) {
	b := NewBoard(15)
	g := NewGameplay(b, Black)
	if g == nil {
		t.Error("want not nil, got nil")
	}
	if g.Board() != b {
		t.Error("want same, got different")
	}
}

func TestGameplayReset(t *testing.T) {
	b := NewBoard(15)
	g := NewGameplay(b, Black)
	g.Move(0, 0)
	want := Black
	got := b.Get(0, 0)
	if got != want {
		t.Errorf("want %d, got %d", want, got)
	}
	g.Reset()
	want = Empty
	got = b.Get(0, 0)
	if got != want {
		t.Errorf("want %d, got %d", want, got)
	}
}

func TestGameplayTurn(t *testing.T) {
	b := NewBoard(15)
	g := NewGameplay(b, Black)
	want := Black
	got := g.Turn()
	if got != want {
		t.Errorf("want %d, got %d", want, got)
	}
	g.Move(0, 0)
	want = White
	got = g.Turn()
	if got != want {
		t.Errorf("want %d, got %d", want, got)
	}
}

func TestGameplayMove(t *testing.T) {
	b := NewBoard(15)
	g := NewGameplay(b, Black)
	want := Black
	got := g.Turn()
	if got != want {
		t.Errorf("want %d, got %d", want, got)
	}
	err := g.Move(0, 0)
	if err != nil {
		t.Errorf("want nil, got %v", err)
	}
	want = White
	got = g.Turn()
	if got != want {
		t.Errorf("want %d, got %d", want, got)
	}
	err = g.Move(0, 0)
	if err == nil {
		t.Errorf("want error, got nil")
	}
	want = White
	got = g.Turn()
	if got != want {
		t.Errorf("want %d, got %d", want, got)
	}
}

func TestGameplayMoveInvalid(t *testing.T) {
	b := NewBoard(15)
	g := NewGameplay(b, Black)
	err := g.Move(-1, 0)
	if err == nil {
		t.Errorf("want error, got nil")
	}
	err = g.Move(0, -1)
	if err == nil {
		t.Errorf("want error, got nil")
	}
	err = g.Move(15, 0)
	if err == nil {
		t.Errorf("want error, got nil")
	}
	err = g.Move(0, 15)
	if err == nil {
		t.Errorf("want error, got nil")
	}
}

func TestGameplayJudge(t *testing.T) {
	b := NewBoard(15)
	g := NewGameplay(b, Black)
	for i := 0; i < 5; i++ {
		err := g.Move(0, i)
		if err != nil {
			t.Errorf("want nil, got %v", err)
		}
		err = g.Move(1, i)
		if i == 4 {
			if err == nil {
				t.Errorf("want error, got nil")
			}
			result := g.Judge()
			if result.Winner != Black {
				t.Errorf("want %d, got %d", Black, result.Winner)
			}
		} else {
			if err != nil {
				t.Errorf("want nil, got %v", err)
			}
		}
	}
}
