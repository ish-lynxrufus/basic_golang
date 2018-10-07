package main

import "fmt"

func main() {
	score := 90
	fmt.Println(score)

	if score > 80 {
		fmt.Println("Great!")
	} else if score > 60 {
		fmt.Println("Good!")
	} else {
		fmt.Println("so so...")
	}

	switch {
	case score > 80:
		fmt.Println("Great!")
	case score > 60:
		fmt.Println("Good!")
	default:
		fmt.Println("so so...")
	}

	for i := 0; i < 10; i++ {
		if i == 5 {
			break
		}
		if i == 3 {
			continue
		}
		fmt.Println(i)
	}

	i := 0
	for i < 10 {
		fmt.Println(i)
		i++
	}

	j := 0
	for {
		fmt.Println(j)
		j++
		if j == 5 {
			break
		}
	}

	s := []int{2, 3, 8}
	for _, v := range s {
		fmt.Println(v)
	}

	m := map[string]int{"tokyo": 13, "kanagawa": 14}
	for key, val := range m {
		fmt.Println(key, val)
	}
}
