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

func main() {
	var x, y, color int
	var t0 = [2]int{1, 0}
	var t45 = [2]int{1, 1}
	var t90 = [2]int{0, 1}
	var t135 = [2]int{-1, 1}
	var t180 = [2]int{-1, 0}
	var t225 = [2]int{-1, 1}
	var t270 = [2]int{0, -1}
	var t315 = [2]int{1, -1}
	var dirs = [8][2]int{t0, t45, t90, t135,
		t180, t225, t270, t315}

	board := initBoard()
	board.printBoard()

	bothSurrender := false
	for !bothSurrender {
		// Wait For Stdin
		for i := 0; i < PLAYER_NUM; i++ {
			if i == 0 {
				fmt.Println("Black Player!")
				color = BLACK
			} else {
				fmt.Println("White Player!")
				color = WHITE
			}
			fmt.Println("Choose X(0~7):")
			fmt.Scanf("%d", &x)
			fmt.Println("Choose Y(0~7):")
			fmt.Scanf("%d", &y)

			board.addStone(x, y, color)
			board.printBoard()
		}
		// if Black.Pass == True && White.Pass == True
		//	bothSurrender := truei
		// break;
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
	B.Coordinates[4][3] = initStone(WHITE)
	B.Coordinates[4][4] = initStone(BLACK)

	return B
}

func (b *Board) addStone(x, y, color int) {
	b.Coordinates[x][y].Color = color
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
	return !isOutOfBounds(x, y)

	stone := b.Coordinates[x][y]

	// If Stone is already Placed
	if stone.Color != EMPTY {
		return false
	}

	//var existDifferentColorDirection []int
	//var neighborStone Stone

	return true
}
