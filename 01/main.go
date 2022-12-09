package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {

	pwd, _ := os.Getwd()
	input, err := os.Open(pwd + "/01/input1.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	fileScanner := bufio.NewScanner(input)
	fileScanner.Split(bufio.ScanLines)

	var cl []int
	var c int
	for fileScanner.Scan() {
		if fileScanner.Text() == "" {
			cl = append(cl, c)
			c = 0
		} else {
			intVal, err := strconv.Atoi(fileScanner.Text())
			if err != nil {
				fmt.Println(err)
				return
			}
			c = c + intVal
		}
	}
	input.Close()

	sort.Ints(cl)

	fmt.Print(cl)
}
