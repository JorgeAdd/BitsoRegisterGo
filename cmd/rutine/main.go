package main

import (
	"fmt"

	"github.com/JorgeAdd/BitsoRegisterGo/BitsoRegisterRutine/internal/controller"
)

func main() {
	fmt.Println("hell rutine")
	controller.GetBitsoInfoController()
}

// func process1(ch chan string) {
// 	for {
// 		time.Sleep(10 * time.Second)
// 		ch <- "hello1"
// 	}
// }

// func main() {
// 	output1 := make(chan string)
// 	go process1(output1)
// loop:
// 	for {
// 		select {
// 		case s1, ok := <-output1:
// 			if !ok {
// 				break loop
// 			}
// 			fmt.Println(s1)
// 		}
// 	}
// }
