package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	pwd, _ := os.Getwd()
	input, err := os.Open(pwd + "/06/input.txt")
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
	var cc int
	for fs.Scan() {
		s := fs.Text()
		var ca [14]string
		for _, c := range s {
			for i, v := range ca {
				if i == 0 {
					continue
				}
				ca[i-1] = v
			}
			ca[13] = string(c)
			cc++
			if cc < 14 {
				continue
			}
			if checkArray(ca) {
				break
			}
		}
	}
	f.Close()

	fmt.Print(cc)
}

func checkArray(a [14]string) bool {
	for i := 0; i < 14; i++ {
		for j := 0; j < 14; j++ {
			if i == j {
				continue
			}
			if a[i] == a[j] {
				return false
			}
		}
	}
	return true
}
