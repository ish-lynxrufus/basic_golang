package main

import "fmt"

func show(t interface{}) {
	// _, ok := t.(japanese)
	// if ok {
	// 	fmt.Println("私は日本人です。")
	// } else {
	// 	fmt.Println("I am not Japanese.")
	// }

	switch t.(type) {
	case japanese:
		fmt.Println("私は日本人です。")
	default:
		fmt.Println("I am not Japanese.")
	}
}

type greeter interface {
	greet()
}

type japanese struct{}
type amarican struct{}

func (j japanese) greet() {
	fmt.Println("こんにちは")
}
func (a amarican) greet() {
	fmt.Println("Hello")
}

func main() {
	greeters := []greeter{japanese{}, amarican{}}
	for _, greeter := range greeters {
		greeter.greet()
		show(greeter)
	}
}
