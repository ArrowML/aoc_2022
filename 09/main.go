package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	DOWN  = "D"
	UP    = "U"
	LEFT  = "L"
	RIGHT = "R"
)

type point struct {
	X int
	Y int
}

func (p *point) Move(d string) {
	switch d {
	case DOWN:
		p.Y--
	case UP:
		p.Y++
	case LEFT:
		p.X--
	case RIGHT:
		p.X++
	}
}

func (p *point) CheckTail(h *point) {
	if h.X > p.X+1 {
		p.X++
		if h.Y > p.Y {
			p.Y++
		}
		if h.Y < p.Y {
			p.Y--
		}
	}
	if h.Y > p.Y+1 {
		p.Y++
		if h.X > p.X {
			p.X++
		}
		if h.X < p.X {
			p.X--
		}
	}
	if h.X < p.X-1 {
		p.X--
		if h.Y > p.Y {
			p.Y++
		}
		if h.Y < p.Y {
			p.Y--
		}
	}
	if h.Y < p.Y-1 {
		p.Y--
		if h.X > p.X {
			p.X++
		}
		if h.X < p.X {
			p.X--
		}
	}
}

func main() {
	pwd, _ := os.Getwd()
	input, err := os.Open(pwd + "/09/input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	part1(input)

	//part2(input)
}

func part1(f *os.File) {
	fs := bufio.NewScanner(f)
	fs.Split(bufio.ScanLines)

	H := point{X: 0, Y: 0}
	T := point{X: 0, Y: 0}
	tp := make(map[point]bool)
	tp[T] = true
	for fs.Scan() {
		line := fs.Text()
		cmd := strings.Split(line, " ")

		dir := cmd[0]
		dis, _ := strconv.Atoi(cmd[1])

		for i := 0; i < dis; i++ {
			H.Move(dir)
			T.CheckTail(&H)
			tp[T] = true
		}

	}
	f.Close()

	fmt.Print(len(tp))
}

func part2(f *os.File) {
	fs := bufio.NewScanner(f)
	fs.Split(bufio.ScanLines)

	var rope [10]*point
	for j := 0; j < 10; j++ {
		rope[j] = &point{X: 0, Y: 0}
	}
	tp := make(map[point]bool)
	tp[*rope[9]] = true
	for fs.Scan() {
		line := fs.Text()
		cmd := strings.Split(line, " ")

		dir := cmd[0]
		dis, _ := strconv.Atoi(cmd[1])

		for i := 0; i < dis; i++ {
			rope[0].Move(dir)
			for r := 1; r < 10; r++ {
				rope[r].CheckTail(rope[r-1])
			}
			tp[*rope[9]] = true
		}

	}
	f.Close()

	fmt.Print(len(tp))

}
