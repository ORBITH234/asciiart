package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("enter 1 argument")
		return
	}
	input := os.Args[1]
	content, err := os.ReadFile("standard.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}
	bannerlines := strings.Split(strings.ReplaceAll(string(content), "\r\n", "\n"), "\n")
	if len(bannerlines) < 760 {
		fmt.Println("Error: banner file has too few lines")
		return
	}
	for row := 0; row < 8; row++ {
		for _, ch := range input {
			index := (int(ch) - 32) * 9
			fmt.Print(bannerlines[index+row])
		}
		fmt.Println()
	}
}
