package main

import "fmt"

func main() {

	const (
		width     = 50
		height    = 10
		emptyCell = " "
		ballCell  = '🥎'
	)

	var (
		px, py int
		cell   rune
	)

	board := make([][]bool, width)
	buf := make([]rune, 0, width*height)

	for row := range board {
		board[row] = make([]bool, height)
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
	fmt.Print(string(buf))

}
