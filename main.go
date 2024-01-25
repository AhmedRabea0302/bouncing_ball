package main

import (
	"fmt"
	"time"
)

func main() {

	const (
		width     = 50
		height    = 10
		emptyCell = " "
		ballCell  = '🥎'
		maxFrames = 1200
	)

	var (
		px, py int
		vx, vy = 1, 1
		cell   rune
	)

	board := make([][]bool, width)
	buf := make([]rune, 0, width*height)

	for row := range board {
		board[row] = make([]bool, height)
	}

	for i := 0; i < maxFrames; i++ {
		px += vx
		py += vy

		if px <= 0 || px >= width-1 {
			vx *= -1
		}

		if py <= 0 || py >= height-1 {
			vy *= -1
		}

		for y := range board[0] {
			for x := range board {
				board[x][y] = false
			}
		}

		board[px][py] = true
		buf = buf[:0]
		for y := range board[0] {
			for x := range board {
				cell = ' '
				if board[x][y] {
					cell = ballCell
				}
				// fmt.Print(string(cell), " ")
				buf = append(buf, cell, ' ')
			}
			buf = append(buf, '\n')
		}
		// screen.MoveTopLeft()
		fmt.Print(string(buf))

		time.Sleep(time.Second / 20)
	}
}
