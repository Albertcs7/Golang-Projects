package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func addcontact(map1 map[string]int) {
	var key, val string
	fmt.Println("ENter key and value:")
	reader := bufio.NewReader(os.Stdin)
	key, _ = reader.ReadString('\n')
	val, _ = reader.ReadString('\n')
	value, _ := strconv.Atoi(val)
	map1[key] = value
}

func viewcontact(map1 map[string]int) {
	fmt.Println("Contact:")
	for index, value := range map1 {
		fmt.Println(index, ".", value)
	}
}

func delcontact(map1 map[string]int) {

}

func main() {
	for {
		var choice int
		fmt.Println("\n1.Add contact\n2.View contact\n3.Delete contact\n4.Exit")
		fmt.Print("Enter a choice: ")
		fmt.Scan(&choice)
		map1 := make(map[string]int)

		switch choice {
		case 1:
			addcontact(map1)
		case 2:
			viewcontact(map1)
		case 3:
			delcontact(map1)
		case 4:
			os.Exit(0)
		default:
			fmt.Println("Enter valid choice...")
		}
	}
}
