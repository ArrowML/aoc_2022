package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var vmap = make(map[string]int)

func main() {

	chars := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

	for i, c := range chars {
		vmap[string(c)] = i + 1
	}

	pwd, _ := os.Getwd()
	input, err := os.Open(pwd + "/03/input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	// part1(input)

	part2(input)
}

func part1(f *os.File) {

	fs := bufio.NewScanner(f)
	fs.Split(bufio.ScanLines)
	var total int
	for fs.Scan() {
		tl := fs.Text()

		ts := strings.Split(tl, "")

		fc := ts[0 : len(ts)/2]
		sc := ts[len(ts)/2 : len(ts)]

		cmap := make(map[string]bool)
		for _, c := range fc {
			cmap[c] = true
		}

		var match string
		for _, c := range sc {
			if _, ok := cmap[c]; ok {
				match = c
				break
			}
		}

		total = total + vmap[match]

	}
	f.Close()

	fmt.Print(total)

}

func part2(f *os.File) {

	fs := bufio.NewScanner(f)
	fs.Split(bufio.ScanLines)
	var total int

	var g []string
	for fs.Scan() {
		g = append(g, fs.Text())
		if len(g) == 3 {

			one := make(map[string]bool)
			for _, c := range g[0] {
				one[string(c)] = true
			}

			two := make(map[string]bool)
			for _, c := range g[1] {
				two[string(c)] = true
			}

			var match string
			for _, c := range g[2] {
				if _, ok := one[string(c)]; ok {
					if _, ko := two[string(c)]; ko {
						match = string(c)
						break
					}
				}
			}
			total = total + vmap[match]
			g = []string{}
		}

	}
	f.Close()

	fmt.Print(total)

}
