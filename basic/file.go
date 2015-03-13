package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	fh, ferr := os.Open("/Users/maxwell/Documents/private.pem")
	if ferr != nil {
		fmt.Printf("An error occurred on opening the inputfile\n" +
			"Does the file exist?\n" +
			"Have you got acces to it?\n")
		return
	}
	defer fh.Close()
	inputread := bufio.NewReader(fh)
	for {
		input, ferr := inputread.ReadString('\n')
		if ferr == io.EOF {
			return
		}
		fmt.Println(strings.TrimSpace(input))
	}
}
