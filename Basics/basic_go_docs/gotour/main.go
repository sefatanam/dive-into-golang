package main

import (
	"fmt"

	"github.com/sefatanam/gotour/src"
)

// func main() {
// 	var yahoo src.Yahoo
// 	rectangle := src.Rectangle{Height: 12.2, Width: 10.8}
// 	yahoo = rectangle
// 	fmt.Println("now rectangle area is : ", rectangle.Area())
// 	src.Scale(rectangle,1.2)
// 	fmt.Println("now rectangle area is : ", rectangle.Area())
// 	fmt.Println("now yahoo area is : ", yahoo.Area())
//
// 	// rectangle.Scale(1.0 / 1.2)
// 	// fmt.Println("now rectangle area is : ", rectangle.Area())
// }

func main() {
	processChan := make(chan int, 1)

	go src.BulkProcess(11, processChan)

	for v:= range processChan{
		fmt.Println("Loading..", v)
	}
}
