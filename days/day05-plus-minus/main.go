package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

/*
 * Complete the 'plusMinus' function below.
 *
 * The function accepts INTEGER_ARRAY arr as parameter.
 */

func plusMinus(arr []int32) {
	var positiveCounter, neagativeCounter, zeroCounter int
	len := len(arr)
	for i := range arr {
		if arr[i] > 0 {
			positiveCounter++
		} else if arr[i] < 0 {
			neagativeCounter++
		} else {
			zeroCounter++
		}
	}
	if len > 0 {
		plusRatio := float32(positiveCounter) / float32(len)
		minusRatio := float32(neagativeCounter) / float32(len)
		zeroRatios := float32(zeroCounter) / float32(len)
		fmt.Printf("%.6f\n", plusRatio)
		fmt.Printf("%.6f\n", minusRatio)
		fmt.Printf("%.6f\n", zeroRatios)
	}
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	nTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	n := int32(nTemp)

	arrTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")

	var arr []int32

	for i := 0; i < int(n); i++ {
		arrItemTemp, err := strconv.ParseInt(arrTemp[i], 10, 64)
		checkError(err)
		arrItem := int32(arrItemTemp)
		arr = append(arr, arrItem)
	}

	plusMinus(arr)
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
