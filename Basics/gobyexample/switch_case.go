package main

import "fmt"

func main() {
	whatAmI := func(i interface{}) { // Infer type from interface
		switch t := i.(type) {
		case bool:
			fmt.Printf("I'm a bool %v\n", t)
		case int:
			fmt.Printf("I'm a int %v\n", t)
		case string:
			fmt.Printf("I'm a string %v", t)
		default:
			fmt.Printf("This is something else %v\n", t)
		}
	}

	whatAmI(true)
	whatAmI(1)
	whatAmI("Sefat Anam")
}
