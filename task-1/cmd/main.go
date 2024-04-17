package main

// Импорты других пакетов
import (
	"fmt"
	_ "github.com/go-loremipsum/loremipsum"
)

// Неявная инициализация пакета
func init() {
	fmt.Println("Hello from init!")
}

// Функция main как точка входа

func main() {
	foo()

	//mess := order.MyOrder()
	//fmt.Println(mess)
}

func foo() {
	fmt.Println("Foo!")
}
