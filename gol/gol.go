package main

func wminus(a, b, max int) int { // minus with wrapping
	res := a - b
	if res < 0 {
		res = max + res + 1
	}
	return res
}

func wplus(a int, b, max int) int { // plus with wrapping
	res := a + b
	if res > max {
		res = res - max - 1
	}
	return res
}

func sumCellNeighbours(col, pos, width, height int, world [][]byte) int {
	ret := byte(0)

	ret += world[col][wplus(pos, 1, height)]  // u
	ret += world[col][wminus(pos, 1, height)] // d
	ret += world[wminus(col, 1, width)][pos]  // l
	ret += world[wplus(col, 1, width)][pos]   // r

	ret += world[wminus(col, 1, width)][wplus(pos, 1, height)]  // ul
	ret += world[wplus(col, 1, width)][wplus(pos, 1, height)]   // ur
	ret += world[wminus(col, 1, width)][wminus(pos, 1, height)] // dl
	ret += world[wplus(col, 1, width)][wminus(pos, 1, height)]  // dr

	return int(ret) / 255
}

func calculateNextState(p golParams, world [][]byte) [][]byte {
	width := p.imageWidth
	height := p.imageHeight
	turns := p.turns

	newworld := world

	for i, col := range world {
		for j, val := range col {
			sum := sumCellNeighbours(i, j, width, height, world)
			if val == 255 { // live
				if sum < 2 {
					newworld[i][j] = 0
				}
			}
		}
	}

	return world
}

func calculateAliveCells(p golParams, world [][]byte) []cell {
	return []cell{}
}
