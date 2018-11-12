package main

import (
	"bufio"
	"fmt"
	"os"
	"os/signal"
	"strconv"
	"strings"
)

var (
	pid = os.Getpid()
)

func main() {
	// make this on different thread
	go func() {
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
					if valueInt >= 10 && valueInt <= 50 {
						panic("Hey something error here")
					}
					fmt.Printf("Square value of %d is %d\n", valueInt, valueInt*valueInt)
				}
			}
		}
		shutdown()
	}()

	// gracefull shutdown
	cDone := make(chan os.Signal, 1)
	signal.Notify(cDone, os.Interrupt)
	<-cDone
	fmt.Println("\nThank You !!")
}

func shutdown() {
	p, err := os.FindProcess(pid)
	if err != nil {
		fmt.Println("err: " + err.Error())
		return
	}

	p.Signal(os.Interrupt)
}
