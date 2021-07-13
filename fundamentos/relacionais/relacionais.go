package main

import (
	"fmt"
	"time"
)

func main() {

	fmt.Println("String", "Banana" == "Banana")
	fmt.Println("!=", 3 != 2)
	fmt.Println("<", 3 < 2)
	fmt.Println(">", 3 > 2)
	fmt.Println("<=", 3 <= 2)
	fmt.Println(">=", 3 >= 2)

	d1 := time.Unix(0, 0)
	d2 := time.Unix(0, 0)

	fmt.Println("Data 1", d2)
	fmt.Println("Data 2:", d2)

	fmt.Println("Datas", d1 == d2)

	type Pessoa struct {
		Nome string
	}

	p1 := Pessoa{"Joao"}
	p2 := Pessoa{"Joao"}

	fmt.Println("Pessoas", p1 == p2)

}