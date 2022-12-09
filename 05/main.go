package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var crates_small = [][]string{{"Z", "N"}, {"M", "C", "D"}, {"P"}}
var crates_big = [][]string{
	{"S", "Z", "P", "D", "L", "B", "F", "C"},
	{"N", "V", "G", "P", "H", "W", "B"},
	{"F", "W", "B", "J", "G"},
	{"G", "J", "N", "F", "L", "W", "C", "S"},
	{"W", "J", "L", "T", "P", "M", "S", "H"},
	{"B", "C", "W", "G", "F", "S"},
	{"H", "T", "P", "M", "Q", "B", "W"},
	{"F", "S", "W", "T"},
	{"N", "C", "R"}}

/*
[C]         [S] [H]
[F] [B]     [C] [S]     [W]
[B] [W]     [W] [M] [S] [B]
[L] [H] [G] [L] [P] [F] [Q]
[D] [P] [J] [F] [T] [G] [M] [T]
[P] [G] [B] [N] [L] [W] [P] [W] [R]
[Z] [V] [W] [J] [J] [C] [T] [S] [C]
[S] [N] [F] [G] [W] [B] [H] [F] [N]
1   2   3   4   5   6   7   8   9

    [D]
[N] [C]
[Z] [M] [P]
*/

func main() {

	pwd, _ := os.Getwd()
	input, err := os.Open(pwd + "/05/input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	//part1(input)

	part2(input)
}

func part1(f *os.File) {

	crates := crates_big

	fs := bufio.NewScanner(f)
	fs.Split(bufio.ScanLines)
	for fs.Scan() {
		line := strings.Replace(fs.Text(), "move ", "", 1)
		line = strings.Replace(line, " from", "", 1)
		line = strings.Replace(line, " to", "", 1)

		ins := strings.Split(line, " ")

		count, _ := strconv.Atoi(ins[0])
		from, _ := strconv.Atoi(ins[1])
		to, _ := strconv.Atoi(ins[2])

		to = to - 1
		from = from - 1

		for i := 0; i < count; i++ {
			crate := crates[from][len(crates[from])-1]
			crates[to] = append(crates[to], crate)
			crates[from] = crates[from][:len(crates[from])-1]
		}
	}
	f.Close()

	for _, c := range crates {
		fmt.Print(c[len(c)-1])
	}

}

func part2(f *os.File) {

	crates := crates_big

	fs := bufio.NewScanner(f)
	fs.Split(bufio.ScanLines)
	for fs.Scan() {
		line := strings.Replace(fs.Text(), "move ", "", 1)
		line = strings.Replace(line, " from", "", 1)
		line = strings.Replace(line, " to", "", 1)

		ins := strings.Split(line, " ")

		count, _ := strconv.Atoi(ins[0])
		from, _ := strconv.Atoi(ins[1])
		to, _ := strconv.Atoi(ins[2])

		to = to - 1
		from = from - 1

		end := len(crates[from])
		start := end - count

		mc := crates[from][start:end]
		crates[to] = append(crates[to], mc...)
		crates[from] = crates[from][:start]
	}
	f.Close()

	for _, c := range crates {
		fmt.Print(c[len(c)-1])
	}
}
