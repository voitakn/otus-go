package main

import (
	"fmt"
	"log"
	"time"
)

func ad() int {
	return 12
}

func adErr() error {
	return fmt.Errorf("My error")
}

func main() {

	for i := 0; i <= 3; i++ {
		fmt.Printf("i_1 = %d\n", i)
	}

	i := 1
	for i <= 4 {
		fmt.Printf("i_2 = %d\n", i)
		i++
	}

	i = 0
	for {
		i++
		if i%2 == 0 {
			continue
		} else if i > 5 {
			break
		}
		fmt.Printf("i_3 = %d\n", i)
	}

	for i := 1; i <= 3; i++ {
		fmt.Printf("i_4 = %d\n", i)
		switch i {
		case 2:
			break
		default:
			fmt.Println("Ok")
		}
		fmt.Println(".")
	}

	letters := []string{"a", "b", "c"}

	for _, letter := range letters {
		fmt.Printf("Index: %d Value:%s\n", i, letter)
	}

	if err := adErr(); err != nil {
		log.Println(err)
	}
}

func main6() {

	//a := 12

	t := time.Now()
	switch {
	case t.Second() < 15 && t.Second() > 2:
		fmt.Printf("<15 %v\n", t)
	case t.Second() < 30:
		fmt.Printf("<30 %v\n", t)
	case t.Second() < 45:
		fmt.Printf("<45 %v\n", t)
	default:
		fmt.Printf(">45 %v\n", t)
	}

	fmt.Println("TEST main 3")
}

func main5() {

	//a := 12

	switch age := ad(); age {
	case 30, 25:
		fmt.Println(30)
		fmt.Println(25)
	case 3 * 4:
		fmt.Println(12)
		fallthrough
	case 2 * ad():
		fmt.Println(24)
	default:
		fmt.Printf("Age = %d\n", age)
	}

	fmt.Println("TEST main 3")
}
func main4() {
	var s any
	s = 12
	switch stype := s.(type) {
	case string:
		fmt.Println("string")
	case bool:
		fmt.Println("boolean")
	case float64:
		fmt.Println("float64")
	case float32:
		fmt.Println("float32")
	case int:
		fmt.Println("int")
	default:
		fmt.Printf("%T", stype)
	}
}

func main3() {
	age, _, _ := 12, "m", "Kaluga"

	//a := 12

	switch age {
	case 30, 25:
		fmt.Println(30)
		fmt.Println(25)
	case 3 * 4:
		fmt.Println(12)
		fallthrough
	case 2 * ad():
		fmt.Println(24)
	default:
		fmt.Printf("Age = %d\n", age)
	}
	fmt.Println("TEST main 3")
}

func main2() {
	age, sex, city := 13, "m", "Kaluga"

	if age <= 14 {
		fmt.Printf("Ребенок\n")
	} else if sex != "m" || city == "Moscow" {
		fmt.Printf("Москвичка\n")
	} else if sex == "m" {
		fmt.Printf("Мужчина\n")
	} else {
		fmt.Printf("Человек %\n", age)
	}
	fmt.Println("TEST main 2")
}

func main1() {

	exist := "active"

	if exist == "disabled" {
		fmt.Printf("Not exist disabled\n")
	} else if exist == "deleted" {
		fmt.Printf("Not exist deleted\n")
	} else {
		fmt.Printf("Exist active\n")
	}

	fmt.Println("TEST main")
}
