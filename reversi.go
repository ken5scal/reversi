package main

import "fmt"

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
			playGround.board[i] = append(playGround.board[i], j)
		}
	}
	fmt.Println(playGround.Score)

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
	board map[int][]int
	board map[int]Coordinate
}

type Coordinate struct {
	Color string
	X     int
	Y     int
}
