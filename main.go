package main

import "fmt"

func main() {
	strSlice := cleanInput("   hello    world   ")
	for subString := range strSlice {
		fmt.Printf("Substring: %v", subString)
	}
}
