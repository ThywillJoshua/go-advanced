package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"time"
)

func sumFile(rd bufio.Reader) int {
	sum := 0

	for {
		line, err := rd.ReadString('\n')
		if err == io.EOF {
			return sum
		}

		if err != nil {
			fmt.Printf("err: %v\n", err)
		}

		num, err := strconv.Atoi(line[:len(line)-1])
		if err != nil {
			fmt.Printf("err: %v\n", err)
		}

		sum += num
	}
}

func main() {
	files := []string{"num1.txt", "num2.txt", "num3.txt", "num4.txt"}
	sum := 0

	for _, file := range files {
		file, err := os.Open(file)
		if err != nil {
			fmt.Printf("err: %v\n", err)
			return
		}

		rd := bufio.NewReader(file)

		calculate := func() {
			fileSum := sumFile(*rd)
			sum += fileSum
		}

		go calculate()
	}

	time.Sleep(2 * time.Second)
	fmt.Println(sum)
}
