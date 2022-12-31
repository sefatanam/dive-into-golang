package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	var fileName string

	fmt.Print("Enter a file name [* must contain extension]:")
	fmt.Scanln(&fileName)
	fmt.Print("Enter your texts: ")
	reader := bufio.NewReader(os.Stdin)
	texts, _ := reader.ReadString('.')
	fmt.Scanln(&texts)
	file, fileError := os.Create(fileName)

	if fileError != nil {
		fmt.Println("Error occurred to create file.", fileError)
		return
	}

	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			fmt.Println("Error occurred to save file.", err)
		}
	}(file)

	_, creationError := file.WriteString(string(texts))

	if creationError != nil {
		fmt.Println("Error occurred to write into file")
	}
	fmt.Println("File is written to", file.Name())

}
