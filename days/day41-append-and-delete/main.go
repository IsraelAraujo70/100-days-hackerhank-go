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
 * Complete the 'appendAndDelete' function below.
 *
 * The function is expected to return a STRING.
 * The function accepts following parameters:
 *  1. STRING s
 *  2. STRING t
 *  3. INTEGER k
 */

func appendAndDelete(s string, t string, k int32) string {
    commonLength := 0
    minLen := len(s)
    if len(t) < minLen {
        minLen = len(t)
    }
    
    for i := 0; i < minLen; i++ {
        if s[i] == t[i] {
            commonLength++
        } else {
            break
        }
    }
    
    deleteOps := len(s) - commonLength  
    appendOps := len(t) - commonLength  
    minOps := deleteOps + appendOps
    
    if minOps > int(k) {
        return "No"
    }

    if minOps == int(k) {
        return "Yes"
    }

    extraOps := int(k) - minOps

    totalOps := len(s) + len(t)
    
    if int(k) >= totalOps {
        return "Yes"
    }

    if extraOps%2 == 0 {
        return "Yes"
    }
    
    return "No"
}

func main() {
    reader := bufio.NewReaderSize(os.Stdin, 16 * 1024 * 1024)

    // Handle OUTPUT_PATH for HackerRank or use stdout for testing
    var writer *bufio.Writer
    if outputPath := os.Getenv("OUTPUT_PATH"); outputPath != "" {
        stdout, err := os.Create(outputPath)
        checkError(err)
        defer stdout.Close()
        writer = bufio.NewWriterSize(stdout, 16 * 1024 * 1024)
    } else {
        writer = bufio.NewWriterSize(os.Stdout, 16 * 1024 * 1024)
    }

    s := readLine(reader)

    t := readLine(reader)

    kTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
    checkError(err)
    k := int32(kTemp)

    result := appendAndDelete(s, t, k)

    fmt.Fprintf(writer, "%s\n", result)

    writer.Flush()
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
