package main

func countBattleships(board [][]byte) int {
	// cells are either X for battleship or "." for empty
	// a battleship has shape 1xk or kx1
	// there are no adjacent ships, i.e. there is always at least one empty cell
	// separating to ships
	// count the number of battleships
	// can do this in one pass by counting the top/left cells
	// of battleships, i.e. a X cell that has no X cell to the left or on top
	bs := byte('X')

	result := 0
	for r, row := range board {
		for c, cell := range row {
			if cell == bs {
				if r > 0 && board[r-1][c] == bs {
					continue
				}
				if c > 0 && board[r][c-1] == bs {
					continue
				}
				result += 1
			}
		}
	}
	return result
}
