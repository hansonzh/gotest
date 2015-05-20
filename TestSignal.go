package main

import (
	"fmt"
	"os"
	"os/signal"
)

func SignalHandle() {
	fmt.Println("handle signal ...")
}

func OnInterrupt(sigHandle func()) {
	signalChan := make(chan os.Signal)
	signal.Notify(signalChan, os.Interrupt)
	go func() {
		for _ = range signalChan {
			sigHandle()
			os.Exit(0)
		}
	}()
}

func main() {
	OnInterrupt(SignalHandle)

	c := make(chan int)
	<-c
}
