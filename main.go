package main

import (
	"fmt"
	"os"
	"os/signal"
	"time"
)

func SetupCloseHandler() {
	c := make(chan os.Signal)
	// signal.Notify(c, syscall.SIGTERM, syscall.SIGHUP, syscall.SIGINT)
	signal.Notify(c)
	go func() {
		s := <-c
		fmt.Println(s)
		fmt.Println("\r- Ctrl+C pressed in Terminal")
		os.Exit(0)
	}()
}

func main() {
	SetupCloseHandler()
	for {
		fmt.Println("Waiting signal...")
		time.Sleep(3 * time.Second)
	}
}
