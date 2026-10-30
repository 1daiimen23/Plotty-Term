package main

import (
	"fmt"
	"math"
)

var h = 10
var w = int(math.Round(2.5 * float64(h)))

type Grid struct {
	grid [][]rune
}

func (g *Grid) Init() {

	g.grid = make([][]rune, h)
	for row := range g.grid {
		g.grid[row] = make([]rune, w)
		for col := range g.grid[row] {
			g.grid[row][col] = ' '
		}
	}
}

func (g *Grid) Draw() {
	for row := range g.grid {
		for col := range g.grid[row] {
			fmt.Printf("%c", g.grid[row][col])
		}
		fmt.Printf("\n")
	}
}
