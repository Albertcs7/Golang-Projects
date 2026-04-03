package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	filename := os.Args[1]
	key := os.Args[2]
	var lineno int
	var file, err = os.Open(filename)
	if err != nil {
		fmt.Println("Error in file loading")
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(strings.ToLower(line), strings.ToLower(key)) {
			fmt.Println("FOund at:", lineno+1, "\nSentence:", line)
		}
		lineno++
	}
}
