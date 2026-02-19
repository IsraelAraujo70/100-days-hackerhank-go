package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	const welcome string = "Wecome to our pizza app!"
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)
	userInput, _ := reader.ReadString('\n')

	numRating, err := strconv.ParseFloat(strings.TrimSpace(userInput), 64)

	if err != nil {
		fmt.Println("The rating should be a number")
	}
	fmt.Println("Thanks for rating our pizza the rating of: ", numRating)
}
