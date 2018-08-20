package main

import (
	"os"
	"fmt"
	"time"
	"os/signal"
	"syscall"
	"sync"
)

func handleSignal(Wg  *sync.WaitGroup) {
	defer Wg.Done()
	ch := make(chan os.Signal, 10)
	signal.Notify(ch, syscall.SIGINT)
	for {
		sig := <-ch
		switch sig {
		case syscall.SIGINT:
			fmt.Println("Exiting, please wait...")
			return
		}
	}
}

func main() {
	Wg := &sync.WaitGroup{}
	Wg.Add(1)
	go handleSignal(Wg)
	time.Sleep(1*time.Second)
	syscall.Kill(syscall.Getpid(), syscall.SIGINT)
	Wg.Wait()
}
