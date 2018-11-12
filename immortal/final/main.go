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
	go immortal(mainFunction)

	// gracefull shutdown
	cDone := make(chan os.Signal, 1)
	signal.Notify(cDone, os.Interrupt)
	<-cDone
	fmt.Println("\nThank You !!")
}

func immortal(fn func()) {
	defer func() {
		if rec := recover(); rec != nil {
			fmt.Println("Panic is taken care", rec, "\nMain function will be invoked again\n")
			immortal(fn)
		}
	}()

	fn()
}

func mainFunction() {
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
}

func shutdown() {
	p, err := os.FindProcess(pid)
	if err != nil {
		fmt.Println("err: " + err.Error())
		return
	}

	p.Signal(os.Interrupt)
}
