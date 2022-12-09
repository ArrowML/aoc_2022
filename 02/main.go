package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type symbol struct {
	name   string
	value  int
	beats  string
	looses string
}

const (
	ROCK     string = "rock"
	PAPER    string = "paper"
	SCISSORS string = "scissors"
	WIN      int    = 6
	DRAW     int    = 3
	LOSS     int    = 0
)

func main() {

	pwd, _ := os.Getwd()
	input, err := os.Open(pwd + "/02/input1.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	part1(input)

	part2(input)

}

func part1(f *os.File) {

	smap := make(map[string]symbol)

	smap["A"] = symbol{name: ROCK, value: 1, beats: SCISSORS}
	smap["B"] = symbol{name: PAPER, value: 2, beats: ROCK}
	smap["C"] = symbol{name: SCISSORS, value: 3, beats: PAPER}

	smap["X"] = symbol{name: ROCK, value: 1, beats: SCISSORS}
	smap["Y"] = symbol{name: PAPER, value: 2, beats: ROCK}
	smap["Z"] = symbol{name: SCISSORS, value: 3, beats: PAPER}

	fs := bufio.NewScanner(f)
	fs.Split(bufio.ScanLines)

	var score int
	for fs.Scan() {
		game := fs.Text()
		plays := strings.Split(game, " ")

		opp := smap[plays[0]]
		me := smap[plays[1]]

		if opp.name == me.name {
			score = score + DRAW + me.value
			continue
		}
		if opp.beats == me.name {
			score = score + LOSS + me.value
			continue
		}
		if me.beats == opp.name {
			score = score + WIN + me.value
		}

	}
	f.Close()

	fmt.Print(score)
}

func part2(f *os.File) {

	smap := make(map[string]symbol)
	smap["A"] = symbol{name: ROCK, beats: SCISSORS, looses: PAPER}
	smap["B"] = symbol{name: PAPER, beats: ROCK, looses: SCISSORS}
	smap["C"] = symbol{name: SCISSORS, beats: PAPER, looses: ROCK}

	vmap := make(map[string]int)
	vmap[ROCK] = 1
	vmap[SCISSORS] = 3
	vmap[PAPER] = 2

	rmap := make(map[string]int)
	rmap["X"] = LOSS
	rmap["Y"] = DRAW
	rmap["Z"] = WIN

	fs := bufio.NewScanner(f)
	fs.Split(bufio.ScanLines)

	var score int
	for fs.Scan() {
		game := fs.Text()
		plays := strings.Split(game, " ")

		opp := smap[plays[0]]
		res := rmap[plays[1]]

		if res == LOSS {
			p := opp.beats
			score = score + vmap[p]
			continue
		}
		if res == DRAW {
			p := opp.name
			score = score + vmap[p] + DRAW
			continue
		}
		if res == WIN {
			p := opp.looses
			score = score + vmap[p] + WIN
		}
	}
	f.Close()

	fmt.Print(score)

}
