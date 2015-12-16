package main

import (
	"fmt"
	"strconv"
)

func main() {
	board := initBoard()
	board.printBoard()

	bothSurrender := false
	for !bothSurrender {
		// Wait For Stdin
		for i := 0; i < numPlayer; i++ {
			// Wait for coordinate input from Stdin

			// Validate Coordinate
			// isValidate()
			// if not:  redo

			// Update Board by flipping Color

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

const (
	EMPTY           = 0
	BLACK           = -1
	WHITE           = 1
	COORDINATE_SIZE = 8
	PLAYER_NUM      = 2
	CELL_LEFT       = "["
	CELL_RIGHT      = "]"
	CELL_EMPTY      = "[ ]"
)

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

func (b Board) addStone(x, y, color int) {
	// b.validate()
	b.Coordinates[x][y].Color = color
	// reverseColor()
}

func (b Board) printBoard() {
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
