package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

type grid struct {
	grid   [][]int
	countX int
	countY int
}

var directions = []string{"left", "right", "down", "up"}

func (g *grid) getPointValue(x, y int) int {
	return g.grid[y][x]
}

func (g *grid) getLine(d string, x, y int) []int {
	var nl []int
	switch d {
	case "right":
		for i := x; i < g.countX-1; i++ {
			p := g.grid[y][i+1]
			nl = append(nl, p)
		}
	case "left":
		for i := x; i > 0; i-- {
			p := g.grid[y][i-1]
			nl = append(nl, p)
		}
	case "down":
		for i := y; i < g.countX-1; i++ {
			p := g.grid[i+1][x]
			nl = append(nl, p)
		}
	case "up":
		for i := y; i > 0; i-- {
			p := g.grid[i-1][x]
			nl = append(nl, p)
		}
	}
	return nl
}

func main() {
	pwd, _ := os.Getwd()
	input, err := os.Open(pwd + "/08/input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	fs := bufio.NewScanner(input)
	fs.Split(bufio.ScanLines)

	var treeGrid grid
	for fs.Scan() {
		line := fs.Text()
		var row []int
		for _, c := range line {
			v, _ := strconv.Atoi(string(c))
			row = append(row, v)
		}
		treeGrid.grid = append(treeGrid.grid, row)
	}
	input.Close()

	part1(&treeGrid)

	// part2(&treeGrid)
}

func part1(g *grid) {
	var vc int
	for y := 0; y <= g.countY-1; y++ {
		if y == 0 || y == g.countY-1 {
			vc = vc + g.countY
			continue
		}
		for x := 0; x <= g.countX-1; x++ {
			if x == 0 || x == g.countX-1 {
				vc++
				continue
			}
			pv := g.getPointValue(x, y)
			vis := false
			for _, d := range directions {
				ts := g.getLine(d, x, y)
				lv := true
				for _, v := range ts {
					if v >= pv {
						lv = false
						break
					}
				}
				if lv {
					vis = true
					break
				}
			}
			if vis {
				vc++
			}
		}
	}

	fmt.Print(vc)
}

func part2(g *grid) {
	var ss []int
	for y := 0; y <= g.countY-1; y++ {
		if y == 0 || y == g.countY-1 {
			continue
		}
		for x := 0; x <= g.countX-1; x++ {
			if x == 0 || x == g.countX-1 {
				continue
			}
			pv := g.getPointValue(x, y)
			var vd [4]int
			for i, d := range directions {
				ts := g.getLine(d, x, y)
				var lv int
				for _, v := range ts {
					lv++
					if v >= pv {
						break
					}
				}
				vd[i] = lv
			}
			s := vd[0] * vd[1] * vd[2] * vd[3]
			ss = append(ss, s)
		}
	}
	var max float64
	for _, s := range ss {
		max = math.Max(float64(max), float64(s))
	}
	fmt.Print(max)
}
