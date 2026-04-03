	package main

	import "fmt"

	func simpleFunction(a, b int) (int,error) {
		var res = a / b
		return res,fmt.Errorf("deniaminator nmust not be zero")
	}

	func main() {
		res,_ := simpleFunction(4, 7)
		fmt.Println(res)
	}
