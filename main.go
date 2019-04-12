package main

import (
	"fmt"
	"go-test-messages/functions"
	"go-test-messages/obj"
	"time"
)

func main() {
	var fact = func(position int8) int64 {
		var result int64 = 1
		for i := 2; i <= int(position); i++ {
			result *= int64(i)
		}
		return result
	}

	fmt.Println(functions.Fib(4))
	fmt.Println(fact(4))

	fmt.Println(obj.Person{Name: "Вася", Contacts: []obj.Contact{
		{Class: "phone", Value: "+7913..."},
		{Class: "email", Value: "email@mail.ru"}}})
}

func Sleep(name string) {
	fmt.Println("sleep", name)
	time.Sleep(300 * time.Millisecond)
}
