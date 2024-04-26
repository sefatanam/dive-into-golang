package textsfile

import (
	"fmt"
	"os"
	"text/template"
)

type Greet struct {
	Name  string
	Place string
}

func TextsMain() {
	rafGreet := Greet{"Chotoo", "Dhaka"}

	filePath := "textsFile/demo.txt"

	file, err := template.New("demo.txt").ParseFiles(filePath)

	if err != nil {
		fmt.Printf("Error occurred: %s\n", err)
		return // Exit the function if there's an error
	}

	err = file.Execute(os.Stdout, rafGreet)

	if err != nil {
		fmt.Printf("Error occurred: %s\n", err)
	}

}
