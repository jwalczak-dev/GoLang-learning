package main

import (
	"fmt"
	"io"
	"os"
)

type readF struct{}

func main() {

	fmt.Println(os.Args[1])

	file, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Println("Error when opening file: ", err)
	}

	//readFile1 := readF{}

	//file implements Reader interface(function Read) so it is possible to use it in io.Copy function
	//io.Copy(readFile1, file)

	io.Copy(os.Stdout, file)

}

func (f readF) Write(p []byte) (n int, err error) {
	fmt.Println(string(p))
	fmt.Println("Just wrote this many bytes:", len(p))
	return len(p), nil
}
