package main

import "fmt"

func main() {

	fmt.Print("Mesma")
	fmt.Print(" linha")

	fmt.Println(" Mesma")
	fmt.Println("linha.")

	x := 3.141516

	xs := fmt.Sprint(x)

	fmt.Println("O valor de x e " + xs)
	fmt.Println("O valor de x e", x)

	fmt.Printf("O valor do x e %.2f.", x)

	a := 1
	b := 1.9999
	c := false
	d := "opa"

	fmt.Printf("\n%d %f %.1f %t %s", a, b, b, c, d)

}
