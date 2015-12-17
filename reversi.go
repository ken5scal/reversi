package main

import (
	"fmt"
	"strconv"
)

const (
	EMPTY               = 0
	BLACK               = -1
	WHITE               = 1
	COORDINATE_SIZE     = 8
	PLAYER_NUM          = 2
	NEIGHBOR_STONE_SIZE = 8
	CELL_LEFT           = "["
	CELL_RIGHT          = "]"
	CELL_EMPTY          = "[ ]"
)

var x, y, color int

func main() {
	board := initBoard()
	board.printBoard()
	player_black := Player{}
	player_black.Color = BLACK
	player_white := Player{}
	player_white.Color = WHITE
	// ポインタとして渡さないとあかん
	var currentPlayer *Player

	for true {
		player_black.Pass = false
		player_white.Pass = false
		for i := 0; i < PLAYER_NUM; i++ {
			if i == 0 {
				fmt.Println("Black Player!")
				color = BLACK
				currentPlayer = &player_black
			} else {
				fmt.Println("White Player!")
				color = WHITE
				currentPlayer = &player_white
			}
			fmt.Println("Choose X(0~7):")
			fmt.Scanf("%d", &x)
			fmt.Println("Choose Y(0~7):")
			fmt.Scanf("%d", &y)

			if board.isValidCoodrinate(x, y, color) {
				board.addStone(x, y, color)
			} else {
				//Validじゃなかったらパス扱い
				fmt.Println("passed")
				(*currentPlayer).Pass = true
			}
			board.printBoard()
		}

		// 両プレイヤーがパスしたら修了
		if player_black.Pass && player_white.Pass {
			fmt.Println("Game Over")
			break
		}
	}
}

type Player struct {
	Color int // Black First
	Pass  bool
}

type Board struct {
	Score       int
	Coordinates [COORDINATE_SIZE][COORDINATE_SIZE]Stone
}

type Stone struct {
	Color int
}

type Direction struct {
	dx int
	dy int
}

func initDirection(dx, dy int) Direction {
	D := Direction{}
	D.dx = dx
	D.dy = dy
	return D
}

func initStone(color int) Stone {
	S := Stone{}
	S.Color = color
	return S
}

func initBoard() Board {
	B := Board{}
	var i, j int

	for i = 0; i < COORDINATE_SIZE; i++ {
		for j = 0; j < COORDINATE_SIZE; j++ {
			B.Coordinates[i][j] = initStone(EMPTY)
		}
	}

	B.Coordinates[3][3] = initStone(WHITE)
	B.Coordinates[3][4] = initStone(BLACK)
	B.Coordinates[4][3] = initStone(BLACK)
	B.Coordinates[4][4] = initStone(WHITE)

	return B
}

func (b *Board) addStone(x, y, color int) {
	b.Coordinates[x][y].Color = color
}

func (b *Board) flipStone(x, y int) {
	stone := &b.Coordinates[x][y]
	if stone.Color == WHITE {
		fmt.Println(x, y, "W2B")
		stone.Color = BLACK
	} else if stone.Color == BLACK {
		fmt.Println(x, y, "B2W")
		stone.Color = WHITE
	}
}

func (b *Board) printBoard() {
	var i, j, color int
	var row_vector string

	// Printing Column Indexes
	fmt.Println("  0  1  2  3  4  5  6  7")

	for i = 0; i < COORDINATE_SIZE; i++ {
		row_vector = ""
		for j = 0; j < COORDINATE_SIZE; j++ {
			color = b.Coordinates[i][j].Color
			switch color {
			case EMPTY:
				row_vector += CELL_EMPTY
			case WHITE:
				row_vector += CELL_LEFT + "W" + CELL_RIGHT
			case BLACK:
				row_vector += CELL_LEFT + "B" + CELL_RIGHT
			}
		}
		fmt.Println(strconv.Itoa(i) + row_vector)
	}
}

func isOutOfBounds(x, y int) bool {
	if x < 0 || x >= COORDINATE_SIZE ||
		y < 0 || y >= COORDINATE_SIZE {
		return true
	}
	return false
}

func (b *Board) isValidCoodrinate(x, y, color int) bool {
	// Out of Bounds
	if isOutOfBounds(x, y) {
		return false
	}

	stone := b.Coordinates[x][y]

	// If Stone is already Placed
	if stone.Color != EMPTY {
		return false
	}

	// Available Direction
	dirs := initDirs()
	for dir, _ := range dirs {
		for i := 1; i < COORDINATE_SIZE; i++ {
			next_x := x + i*dir.dx
			next_y := y + i*dir.dy
			if isOutOfBounds(next_x, next_y) {
				break
			} else if b.Coordinates[next_x][next_y].Color == EMPTY {
				break
			} else if color == b.Coordinates[next_x][next_y].Color {
				for j := i - 1; j >= 1; j-- {
					new_new_x := x + j*dir.dx
					new_new_y := y + j*dir.dy
					fmt.Println(dir, new_new_x, new_new_y, "Execute flipping")
					b.flipStone(new_new_x, new_new_y)
				}
				break
			} else if color != b.Coordinates[next_x][next_y].Color {
				dirs[dir] = true
			}
		}
	}

	for _, boolean := range dirs {
		if boolean {
			return true
		}
	}

	return false
}

func initDirs() map[Direction]bool {
	dirs := make(map[Direction]bool)

	t0 := initDirection(1, 0)
	t45 := initDirection(1, 1)
	t90 := initDirection(0, 1)
	t135 := initDirection(-1, 1)
	t180 := initDirection(-1, 0)
	t225 := initDirection(-1, 1)
	t270 := initDirection(0, -1)
	t315 := initDirection(1, -1)
	Ds := [8]Direction{t0, t45, t90, t135, t180, t225, t270, t315}
	for d := range Ds {
		dirs[Ds[d]] = false //
	}
	return dirs
}
