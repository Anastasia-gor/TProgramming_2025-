package main

import (
	"fmt"
	"math"
)

func main() {
	// Данные из задачи А
	a := 0.4
	b := 0.8
	x_n := 3.2

	// Данные из задачи Б
	x_values := []float64{4.48, 3.56, 2.78, 5.28, 3.21}

	// Вычисление функции Y для каждого значения x
	for _, x := range x_values {
		y := calculateY(a, b, x)
		fmt.Printf("Для x = %.2f, Y = %.2f\n", x, y)
	}
}

// Функция для вычисления Y
func calculateY(a, b, x float64) float64 {
	numerator := math.Pow(a, x) - math.Pow(b, x)
	denominator := math.Log(a / b)
	cubeRoot := math.Pow(a*b, 1/3)

	return numerator / denominator * cubeRoot
}
