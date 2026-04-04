package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func addcontact(map1 map[string]int, reader *bufio.Reader) {
	var key, val string
	fmt.Println("Enter Name:")
	key, _ = reader.ReadString('\n')
	key = strings.TrimSpace(key)

	fmt.Println("Enter Phone Number:")
	val, _ = reader.ReadString('\n')
	val = strings.TrimSpace(val)

	value, err := strconv.Atoi(val)
	if err != nil {
		fmt.Println("Invalid number!")
		return
	}

	map1[key] = value
	fmt.Println("Contact added successfully!")
}

func viewcontact(map1 map[string]int) {
	fmt.Println("Contact:")
	for index, value := range map1 {
		fmt.Println(index, "->", value)
	}
}

func delcontact(map1 map[string]int, reader *bufio.Reader) {
	fmt.Println("Enter name to delete:")
	del, _ := reader.ReadString('\n')
	del = strings.TrimSpace(del)
	delete(map1, del)
	fmt.Println(del, "deleted")
}

func main() {
	map1 := make(map[string]int)
	reader := bufio.NewReader(os.Stdin)
	for {

		fmt.Println("\n1.Add contact\n2.View contact\n3.Delete contact\n4.Exit")
		fmt.Print("Enter a choice: ")
		inp, _ := reader.ReadString('\n')
		inp = strings.TrimSpace(inp)
		choice, _ := strconv.Atoi(inp)
		switch choice {
		case 1:
			addcontact(map1, reader)
		case 2:
			viewcontact(map1)
		case 3:
			delcontact(map1, reader)
		case 4:
			os.Exit(0)
		default:
			fmt.Println("Enter valid choice...")
		}
	}
}
