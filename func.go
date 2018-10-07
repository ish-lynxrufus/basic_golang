package main

import "fmt"

func hi(name string) (message string) {
	message = "hi! " + name
	return
}

func swap(a, b int) (int, int) {
	return b, a
}

func main() {
	name := "lynxrufus"
	fmt.Println(hi(name))

	fmt.Println(swap(2, 5))

	f := func(a, b int) (int, int) {
		return b, a
	}
	fmt.Println(f(4, 6))

	func(message string) {
		fmt.Println(message)
	}("lynxrufus")
}
