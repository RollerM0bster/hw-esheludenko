package main

import (
	"context"
	"flag"
	"fmt"
	"net"
	"os"
	"os/signal"
	"time"
)

func main() {
	var timeout time.Duration
	flag.DurationVar(&timeout, "timeout", time.Second*10, "timeout")
	flag.Parse()

	if len(flag.Args()) < 2 {
		fmt.Println("Usage: go-telnet [--timeout=10s] host port")
		os.Exit(1)
	}

	address := net.JoinHostPort(flag.Args()[0], flag.Args()[1])
	client := NewTelnetClient(address, timeout, os.Stdin, os.Stdout)
	if err := client.Connect(); err != nil {
		fmt.Printf("Connect Error: %v\n", err)
		os.Exit(1)
	}
	defer client.Close()
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	errCh := make(chan error, 2)
	go func() {
		if err := client.Send(); err != nil {
			errCh <- fmt.Errorf("send Error: %w", err)
		}
	}()

	go func() {
		if err := client.Receive(); err != nil {
			errCh <- fmt.Errorf("receive Error: %w", err)
		}
	}()

	select {
	case err := <-errCh:
		fmt.Printf("Error: %v\n", err)
		stop()
	case <-ctx.Done():
		fmt.Printf("Shutting down\n")
	}
}
