package main

import "fmt"

// program entry point
func main() {
	str := "Hello Stuttgart Gohpers !"
	for _, s := range str {
		fmt.Printf("%c", s)
	}
	fmt.Println()
}
