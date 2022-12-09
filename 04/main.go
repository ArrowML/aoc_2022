package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	pwd, _ := os.Getwd()
	input, err := os.Open(pwd + "/04/input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	//part1(input)

	part2(input)
}

func part1(f *os.File) {
	fs := bufio.NewScanner(f)
	fs.Split(bufio.ScanLines)
	var count int
	for fs.Scan() {
		rl := fs.Text()
		rs := strings.Split(rl, ",")

		ro := strings.Split(rs[0], "-")
		rt := strings.Split(rs[1], "-")

		ros, _ := strconv.Atoi(ro[0])
		roe, _ := strconv.Atoi(ro[1])

		rts, _ := strconv.Atoi(rt[0])
		rte, _ := strconv.Atoi(rt[1])

		if ros <= rts && roe >= rte {
			count = count + 1
			continue
		}

		if rts <= ros && rte >= roe {
			count = count + 1
		}

	}
	f.Close()

	fmt.Print(count)
}

func part2(f *os.File) {
	fs := bufio.NewScanner(f)
	fs.Split(bufio.ScanLines)
	var count int
	for fs.Scan() {
		rl := fs.Text()
		rs := strings.Split(rl, ",")

		ro := strings.Split(rs[0], "-")
		rt := strings.Split(rs[1], "-")

		ros, _ := strconv.Atoi(ro[0])
		roe, _ := strconv.Atoi(ro[1])

		rts, _ := strconv.Atoi(rt[0])
		rte, _ := strconv.Atoi(rt[1])

		if roe >= rts && ros <= rte {
			count = count + 1
			continue
		}

		if rte >= ros && rts <= roe {
			count = count + 1
			continue
		}

	}
	f.Close()

	fmt.Print(count)

}
