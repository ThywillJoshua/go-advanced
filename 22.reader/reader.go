package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func main() {
	r := bufio.NewReader(os.Stdin)

	var sum int

	for {
		input, inputErr := r.ReadString(' ')
		n := strings.TrimSpace(input)

		if n == "" {
			continue
		}

		num, err := strconv.Atoi(n)
		if err != nil {
			fmt.Println(err)
		} else {
			sum += num
		}

		if inputErr == io.EOF {
			break
		}

		if inputErr != nil {
			fmt.Println("Error reading Stdin:", inputErr)
			return
		}

	}

	fmt.Printf("Sum: %v", sum)
}
