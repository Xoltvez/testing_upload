package main

import (
	"fmt"
	"math/rand"
	"time"
)

func getToppingCount() int {
	var total int

	fmt.Print("Banyak Topping: ")
	fmt.Scanln(&total)

	return total
}

func simulateToppings(total int) int {
	rand.Seed(time.Now().UnixNano())
	countInside := 0

	for i := 0; i < total; i++ {
		x := rand.Float64()
		y := rand.Float64()

		if (x-0.5)*(x-0.5)+(y-0.5)*(y-0.5) <= 0.25 {
			countInside++
		}
	}
	return countInside
}

func estimatePi(total, inside int) float64 {
	prop := float64(inside) / float64(total)

	return prop * 4
}

func showResult(total, inside int, pi float64) {
	fmt.Printf("Topping pada Pizza: %d\n", inside)
	fmt.Printf("PI : %.10f\n", pi)
}

func main() {
	topping := getToppingCount()
	jatuhDiPizza := simulateToppings(topping)
	pi := estimatePi(topping, jatuhDiPizza)

	showResult(topping, jatuhDiPizza, pi)
}
