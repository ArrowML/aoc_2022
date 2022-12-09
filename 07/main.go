package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type file struct {
	name string
	size int
}

type directory struct {
	dirs   map[string]*directory
	files  []file
	parent *directory
	size   int
}

func main() {
	pwd, _ := os.Getwd()
	input, err := os.Open(pwd + "/07/input.txt")
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

	root := directory{
		dirs:   make(map[string]*directory),
		files:  []file{},
		parent: nil,
		size:   0,
	}
	currentDir := &root
	for fs.Scan() {
		line := fs.Text()
		cmd := strings.Split(line, " ")
		if cmd[0] == "$" {
			if cmd[1] == "cd" {
				if cmd[2] == "/" {
					continue
				}
				if cmd[2] == ".." {
					currentDir = currentDir.parent
					continue
				}
				currentDir = currentDir.dirs[cmd[2]]
				continue
			}
			if cmd[1] == "ls" {
				continue
			}
		}
		if cmd[0] == "dir" {
			n := cmd[1]
			d := directory{
				dirs:   make(map[string]*directory),
				files:  []file{},
				parent: currentDir,
				size:   0,
			}
			currentDir.dirs[n] = &d
			continue
		}

		s, _ := strconv.Atoi(cmd[0])
		f := file{
			size: s,
			name: cmd[1],
		}
		currentDir.files = append(currentDir.files, f)
		currentDir.size = currentDir.size + s
	}
	f.Close()

	// Part 1
	var total int
	calcDirSizes(&root, &total)
	unused := 70000000 - root.size
	diff := 30000000 - unused

	var dl []directory
	calcMinDirSize(&root, diff, &dl)

	for _, d := range dl {
		fmt.Printf("%s \n", d.size)
	}

}

func calcDirSizes(dir *directory, t *int) int {
	for _, d := range dir.dirs {
		dir.size = dir.size + calcDirSizes(d, t)
	}
	if dir.size <= 100000 {
		*t = *t + dir.size
	}
	return dir.size
}

func calcMinDirSize(dir *directory, min int, dl *[]directory) {
	for _, d := range dir.dirs {
		calcMinDirSize(d, min, dl)
	}
	if dir.size >= min {
		*dl = append(*dl, *dir)
	}
}
