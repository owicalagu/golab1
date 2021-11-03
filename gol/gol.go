package main

func updateCell(world [][]byte, x, y int, p golParams) byte {
	alive := 0
	for i := -1; i <= 1; i++ {
		for j := -1; j <= 1; j++ {
			if (j != 0 || i != 0) && world[(y+j+p.imageHeight)%p.imageHeight][(x+i+p.imageWidth)%p.imageWidth] == 255 {
				alive++
			}
		}
	}

	if alive == 3 || alive == 2 && world[y][x] == 255 {
		return 255
	} else {
		return 0
	}
}

func calculateNextState(p golParams, world [][]byte) [][]byte {
	newState := make([][]byte, p.imageHeight)
	for i := range newState {
		newState[i] = make([]byte, p.imageWidth)
	}
	for y := 0; y < p.imageHeight; y++ {
		for x := 0; x < p.imageWidth; x++ {
			newState[y][x] = updateCell(world, x, y, p)
		}
	}

	return newState
}

func calculateAliveCells(p golParams, world [][]byte) []cell {
	var cells []cell

	for i, col := range world {
		for j, val := range col {
			if val != 0 {
				cells = append(cells, cell{j, i})
			}
		}
	}
	return cells
}
