// Boiling prints the boiling point of water
package main

import "fmt"

const boilingF = 212.0

func main() {
	var f = boilingF
	var c = (f - 32) * 5 / 9
	fmt.Printf("the boiling point of water = %g degrees Fahrenheit or %g degrees Celsius\n", f, c)
	// Output:
	// boiling point = 212 degrees Fahrenheight or 100 degrees Celsius
}
