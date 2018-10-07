package main

import "fmt"

type prefecture struct {
	id    int
	name  string
	score int
}

func (p prefecture) show() {
	fmt.Printf("id:%d, name:%s, score:%d\n", p.id, p.name, p.score)
}

func (p *prefecture) up() {
	p.score++
}

func main() {
	// kanagawa := new(prefecture)
	// kanagawa.id = 14
	// kanagawa.name = "kanagawa"

	kanagawa := prefecture{id: 14, name: "kanagawa", score: 50}
	kanagawa.up()
	kanagawa.show()

	// fmt.Println(kanagawa)
}
