package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Hello there")
	fmt.Print("Entered Text: ")
	input, _ := reader.ReadString('\n')
	fmt.Print(input)
}
