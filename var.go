package main

import "fmt"

func main() {
	var str string
	var intger int
	float := 12.3
	var boolean bool
	fmt.Printf("string: %s, int: %d, float: %f, bool: %t\n", str, intger, float, boolean)

	const name = "ish"
	// name = "lynxrufus"
	fmt.Println(name)

	const (
		sun = iota
		mon
		tue
		wed
		thu
		fri
		sat
	)
	fmt.Println(sun, mon, tue, wed, thu, fri, sat)

	a := "hello "
	b := "world"
	s := a + b
	fmt.Println(s)

	t := true
	f := false
	fmt.Println(t && f)
	fmt.Println(t || f)
	fmt.Println(!t)

	v := 5
	var pv *int
	pv = &v
	fmt.Println(v, pv, *pv)
}
