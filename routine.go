package main

import (
	"fmt"
	"time"
)

func main_task(result chan string) {
	time.Sleep(time.Second * 2)
	fmt.Println("main_task finished!")
	result <- "main_task result"
}

func sub_task() {
	fmt.Println("sub_task finished!")
}

func main() {
	result := make(chan string)

	go main_task(result)
	go sub_task()

	fmt.Println(<-result)

	// time.Sleep(time.Second * 3)
}
