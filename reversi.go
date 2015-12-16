package main

import (
	"fmt"
	"strconv"
)

func main() {
	//fmt.Println("Hellog world")
	// Number of Player
	numPlayer := 2

	// Initialize Players
	black := Player{}
	black.Color = 0
	white := Player{}
	white.Color = 1
	fmt.Println(black.Color, white.Color)

	// Who will be the first one?

	//initialize 8x8board
	playGround := Board{}
	for i := 0; i < 8; i++ {
		for j := 0; j < 8; j++ {
			//playGround.board[i] = append(playGround.board[i], j)

		}
	}
	fmt.Println(playGround.Score)
	fmt.Println("  0  1  2  3  4  5  6  7")
	fmt.Println("0[ ][ ][ ][ ][ ][ ][ ][ ]")
	fmt.Println("1[ ][ ][ ][ ][ ][ ][ ][ ]")
	fmt.Println("2[ ][ ][ ][ ][ ][ ][ ][ ]")
	fmt.Println("3[ ][ ][ ][ ][ ][ ][ ][ ]")
	fmt.Println("4[ ][ ][ ][ ][ ][ ][ ][ ]")
	fmt.Println("5[ ][ ][ ][ ][ ][ ][ ][ ]")
	fmt.Println("6[ ][ ][ ][ ][ ][ ][ ][ ]")
	fmt.Println("7[ ][ ][ ][ ][ ][ ][ ][ ]")

	// Put First 4 Stones

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
	Score int
	board [COORDINATE_SIZE][COORDINATE_SIZE]Stone
}

type Stone struct {
	Color int
}

const (
	EMPTY           = 0
	BLACK           = 1
	WHITE           = 2
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
			B.board[i][j] = initStone(EMPTY)
		}
	}

	B.board[3][3] = initStone(WHITE)
	B.board[3][4] = initStone(BLACK)
	B.board[4][3] = initStone(WHITE)
	B.board[4][4] = initStone(BLACK)

	return B
}

func (b Board) printBoard() {
	var i, j, color int
	var row_vector string

	// Printing Column Indexes
	fmt.Println("  0  1  2  3  4  5  6  7")

	for i = 0; i < COORDINATE_SIZE; i++ {
		for j = 0; j < COORDINATE_SIZE; j++ {
			color = b.board[i][j].Color
			row_vector = ""
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
