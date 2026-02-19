package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	const welcome string = "Hello user!"
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the rating of our pizza")

	// comma ok || err ok
	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Try again later")
		return
	}
	rating, err := strconv.Atoi(input)
	if err != nil {
		fmt.Println("The input must be an int")
		return
	}

	fmt.Println("Thx 4 the rating of: ", rating)
}
