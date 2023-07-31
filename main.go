package main

import "fmt"

func main() {
	fmt.Print("Enter temperature in Celsium:")
	var input float64
	fmt.Scanf("%f", &input)
	convertTemp(input)
}
func convertTemp(input float64) {
	output := (input - 32) * 5 / 9
	fmt.Print(output)
}
