package main

// y = x*x x = [-5, 5] и y = [-5, 5]

// значение x в столбце c:xStart + (xEnd - xStart) / (w - 1) * c., где w это количество столбцов, а c номер столбца
// значение y в строке c: (yStart - yEnd) / (w - 1) * c

func main() {
	grid := Grid{}
	grid.Init()
	grid.Draw()
}
