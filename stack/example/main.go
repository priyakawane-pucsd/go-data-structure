package main

import (
	"fmt"
	"go-data-structure/stack"
	"log"
)

func main() {
	stack := stack.New(2)
	for i := 0; i < 5; i++ {
		error := stack.Push(i + 1)
		if error != nil {
			log.Print(error)
			break
		}
		log.Print(stack)
	}
	fmt.Println("The value popped from the stack is given as:", stack)

	for i := 0; i < 5; i++ {
		if !stack.IsEmpty() {
			x, y := stack.Pop()
			if y == true {
				fmt.Println("Pop element of stack is>>>>>", x)
			}
		} else {
			fmt.Println("stack is empty")
			break
		}
	}
}
