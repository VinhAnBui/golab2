package main

import (
	"fmt"
	"log"
	"os"
	"runtime/trace"
	"time"
)

func foo(channel chan string) {
	for {
		fmt.Println("\nFoo is sending: ping")
		channel <- "ping"

		message := <-channel
		fmt.Println("Foo has received:", message)
	}

}

func bar(channel chan string) {
	for {
		message := <-channel
		fmt.Println("bar has received:", message)

		fmt.Println("\nbar is sending: pong")
		channel <- "pong"
	}
}

func pingPong() {
	c := make(chan string)
	go foo(c) // Nil is similar to null. Sending or receiving from a nil chan blocks forever.
	go bar(c)
	time.Sleep(500 * time.Millisecond)
}

func main() {
	f, err := os.Create("trace.out")
	if err != nil {
		log.Fatalf("failed to create trace output file: %v", err)
	}
	defer func() {
		if err := f.Close(); err != nil {
			log.Fatalf("failed to close trace file: %v", err)
		}
	}()

	if err := trace.Start(f); err != nil {
		log.Fatalf("failed to start trace: %v", err)
	}
	defer trace.Stop()

	pingPong()
}
