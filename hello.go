// コメント

/*
コメント
コメント
*/

package main

import "fmt"

func main() {
	// var message string
	// message = "hello world"
	// var message string = "hello world"
	// var message = "hello world"
	message := "hello world"
	fmt.Println(message)

	// var a, b int
	// a, b = 10, 15
	a, b := 10, 15

	fmt.Println(a)
	fmt.Println(b)

	// var (
	// 	c int
	// 	d string
	// )
	// c = 20
	// d = "hoge"
	var (
		c = 20
		d = "hoge"
	)

	fmt.Println(c)
	fmt.Println(d)
}
