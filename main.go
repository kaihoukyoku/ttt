// Copyright 2017 kaihoukyoku <kaihoukyoku@cock.li>
// This source code is licensed under the zlib license.
// See the LICENSE file for more information.

package main

import (
	"fmt"
	"os"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	width = 800
	height = 600
	title = "TTT"
)

const (
	Empty = iota
	X
	O
	Draw
)

var (
	board [3][3]int
	turn int
	row, col int
	mx, my int32
	turnip string //TODO: get rid of this
	winner int
)

func initialize() {
	for i := range board {
		for j := range board[i] {
			board[i][j] = Empty
		}
	}
	turn = X
	winner = Empty
}

func update() {
	updateMouse()
	mouseToBoard()
	checkWin()
	processInput()
}

func draw() {
	rl.BeginDrawing()

	rl.ClearBackground(rl.RayWhite)

	drawBoard()
	drawHUD()

	rl.EndDrawing()
}

func doTurnip(x int) {
	if x == O {
		turnip = "O"
	} else if x == X {
		turnip = "X"
	} else if x == Draw {
		turnip = "DRAW"
	} else {
		turnip = ""
	}
}

func updateMouse() {
	mx, my = rl.GetMouseX(), rl.GetMouseY()
}

func mouseToBoard() {
	switch {
	case mx < width/3:
		row = 0
	case mx > width/3*2:
		row = 2
	default:
		row = 1
	}
	switch {
	case my < height/3:
		col = 0
	case my > height/3*2:
		col = 2
	default:
		col = 1
	}
}

func processInput() {
	if winner == Empty {	
		if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
			if board[row][col] == Empty {
				board[row][col] = turn			
				if turn == O {
					turn = X
				} else {
					turn = O
				}
			}
		}
	}

	if rl.IsKeyPressed(rl.KeyR) {
		initialize()
		return
	}
	if rl.IsKeyPressed(rl.KeyQ) {
		os.Exit(0)
		return
	}
}

func checkWin() { //TODO: fix this shit holy fuck
	isFull := true
	for i := range board {
		for j := range board[i] {
			if board[i][j] == Empty {
				isFull = false
			}
		}
	}
	if isFull {
		winner = Draw
		return
	}
	for i := 0; i < 3; i++ {
		if board[i][0] != Empty && board[i][0] == board [i][1] && board[i][1] == board[i][2] {
			winner = board[i][0]
			return
		}
		if board[i][0] != Empty && board[0][i] == board[1][i] && board[1][i] == board[2][i] {
			winner = board[0][i]
			return
		}
	}
	if board[0][0] != Empty && board[0][0] == board[1][1] && board[1][1] == board[2][2] {
		winner = board[0][0]
		return
	}
	if board[0][2] != Empty && board[0][2] == board[1][1] && board[1][1] == board[2][0] {
		winner = board[0][2]
		return
	}
}

func drawBoard() {
	rl.DrawRectangle(width/3, height/8, 3, height/4*3, rl.Black)
	rl.DrawRectangle(width/3*2, height/8, 3, height/4*3, rl.Black)
	rl.DrawRectangle(width/8, height/3, width/4*3, 3, rl.Black)
	rl.DrawRectangle(width/8, height/3*2, width/4*3, 3, rl.Black)

	for i := range board {
		for j := range board[i] {
			doTurnip(board[i][j])
			rl.DrawText(turnip, int32(width/3*(i+1)-width/6)-35, int32(height/3*(j+1)-height/6)-50, 120, rl.Maroon)
		}
	}

	if winner == Empty && board[row][col] == Empty {
		doTurnip(turn)
		rl.DrawText(turnip, int32(width/3*(row+1)-width/6)-35, int32(height/3*(col+1)-height/6)-50, 120, rl.Fade(rl.Maroon, 0.3))
	}
}

func drawHUD() {
	rl.DrawText("RESTART: R", 10, height-25, 20, rl.Maroon)
	rl.DrawText("QUIT: Q", width-85, height-25, 20, rl.Maroon)
	doTurnip(turn)
	if winner != Empty {
		doTurnip(winner)
		rl.DrawText(fmt.Sprintf("WINNER: %v", turnip), 10, 10, 20, rl.Maroon)
	} else {
	rl.DrawText(fmt.Sprintf("TURN: %v", turnip), 10, 10, 20, rl.Maroon)

	}
}

func main() {
	rl.InitWindow(width, height, title)
	defer rl.CloseWindow()
	rl.SetTargetFPS(60)

	initialize()
	for !rl.WindowShouldClose() {
		update()
		draw()
	}
}