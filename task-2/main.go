package main

import "fmt"

var n = false
var MyN = false

func main() {
	var n = 25
	if true {
		n, err := fmt.Println("Hello World") // Hello World
		if err != nil {
			fmt.Println("Error: ", n, err)
		}
		fmt.Println("Success:", n) // Success: 12
	}
	fmt.Println(&n) // 0. Почему?
	main2("true")
}

func main2(inf any) {
	x := 15           // Тип int
	xPtr := &x        // Тип *int
	var pages int = 0 // Тип *int, значение nil
	var page *string
	//var inf any // = "23"
	sb := []byte(inf.(string))
	fmt.Println(sb)

	fmt.Printf("%T %+v \n", inf, inf)
	inf = "23"
	//fmt.Printf("%T %+v \n", inf, inf.(bool))

	if page == nil {
		fmt.Printf("%T %+v \n", pages, pages)
	}

	fmt.Printf("%T %+v \n", xPtr, xPtr)

	x = 100
	fmt.Printf("%T %+v \n", *xPtr, *xPtr)
}
