package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	Greet(os.Stdout, "Elodie")
}

func Greet(writter io.Writer, name string) {
	fmt.Fprintf(writter, "Hello, %s", name)
}

func print(str string) {
	fmt.Printf("%s", str)
}
