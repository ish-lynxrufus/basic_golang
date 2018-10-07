package main

import "fmt"

func main() {
	var a [5]int
	a[0] = 3
	a[2] = 5
	a[4] = 10
	// fmt.Println(a)

	b := [...]int{1, 3, 5}
	fmt.Println(b)
	fmt.Println(len(b))

	s := a[2:4]
	s[1] = 20
	fmt.Println(a)
	fmt.Println(s)
	fmt.Println(len(s))
	fmt.Println(cap(s))

	slice := []int{1, 3, 5}
	slice = append(slice, 8, 2, 10)
	cop := make([]int, len(slice))
	n := copy(cop, slice)

	fmt.Println(slice)
	fmt.Println(cop)
	fmt.Println(n)

	// m := make(map[string]int)
	// m["tokyo"] = 13
	// m["kanagawa"] = 14
	m := map[string]int{"tokyo": 13, "kanagawa": 14}
	fmt.Println(m)
	fmt.Println(len(m))
	delete(m, "tokyo")
	fmt.Println(m)

	v, ok := m["tokyo"]
	fmt.Println(v)
	fmt.Println(ok)
}
