package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("Enter a value (blank to exit) : ")
		valueStr, err := reader.ReadString('\n')
		valueStr = strings.TrimSuffix(valueStr, "\n")
		if err != nil {
			fmt.Println("read error: " + err.Error())
		} else {
			if valueStr == "" {
				break
			}
			valueInt, err := strconv.Atoi(valueStr)
			if err != nil {
				fmt.Println("cast error: " + err.Error())
			} else {
				fmt.Printf("Square value of %d is %d\n", valueInt, valueInt*valueInt)
			}
		}
	}
	fmt.Println("Thank You !!")
}
