package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("Hello World")
	v := "sql"
	switch v {
	case "sql":
		fmt.Printf("Connecting to sql")
	case "nosql":
		fmt.Printf("connection nosql")
	default:
		fmt.Printf("no connection")
	}
	x, n := 2.0, 2.0
	if u := math.Pow(x, n); u < 2 {
		fmt.Printf("less than 2 %.lf", u)
	} else {
		fmt.Printf("%v greater than 2", u)
	}
}
