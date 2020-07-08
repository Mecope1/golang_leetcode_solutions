package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

// Complete the sockMerchant function below.
func sockMerchant(n int32, ar []int32) int32 {
	sock_type := []int32{ar[0]}
	sock_count := []int32{1}

	for _, val := range ar[1:] {
		fmt.Println("val:", val)
		for type_ind, type_val := range sock_type {

			fmt.Println("type_val:", val)

			if val == type_val {
				sock_count[type_ind] += 1
				break
			} else if type_ind == len(sock_type) - 1 {
				fmt.Println("appending new sock type!", val)
				sock_type = append(sock_type, val)
				sock_count = append(sock_count, 1)
			}
		}
	}

	var num_pairs int32 = 0

	for _, val := range sock_count {
		num_pairs += val / 2
	}

	//fmt.Println("type array: ", sock_type)
	//fmt.Println("count array: ", sock_count)
	//fmt.Println(num_pairs)
	return num_pairs
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024 * 1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 1024 * 1024)

	nTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)
	n := int32(nTemp)

	arTemp := strings.Split(readLine(reader), " ")

	var ar []int32

	for i := 0; i < int(n); i++ {
		arItemTemp, err := strconv.ParseInt(arTemp[i], 10, 64)
		checkError(err)
		arItem := int32(arItemTemp)
		ar = append(ar, arItem)
	}

	result := sockMerchant(n, ar)

	fmt.Fprintf(writer, "%d\n", result)

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

