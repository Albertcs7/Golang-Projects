package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func addTask(arr []string, count int) int {
	reader := bufio.NewReader(os.Stdin)

	reader.ReadString('\n') 

	fmt.Println("Enter the task:")
	task, _ := reader.ReadString('\n')
	task = strings.TrimSpace(task)

	if count < len(arr) {
		arr[count] = task
		count++
	} else {
		fmt.Println("Task list is full")
	}

	return count
}

func delTask(arr []string, count int) int {
	var tno int
	fmt.Println("Enter the task number:")
	fmt.Scan(&tno)

	if tno < 1 || tno > count {
		fmt.Println("Invalid task number")
		return count
	}
	for i := tno - 1; i < count-1; i++ {
		arr[i] = arr[i+1]
	}

	count--
	return count
}

func viewTask(arr []string, count int) {
	if count == 0 {
		fmt.Println("No tasks available")
		return
	}
	fmt.Println("\nTasks")
	for i := 0; i < count; i++ {
		fmt.Println(i+1, ".", arr[i])
	}
}

func main() {
	count := 0
	var choice int
	arr := make([]string, 10)

	for {
		fmt.Println("\n1.Add Task\n2.View Task\n3.Delete Task\n4.Exit")
		fmt.Print("Enter a choice: ")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			count = addTask(arr, count)
		case 2:
			viewTask(arr, count)
		case 3:
			count = delTask(arr, count)
		case 4:
			os.Exit(0)
		default:
			fmt.Println("Enter valid choice...")
		}
	}
}
