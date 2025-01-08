package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"net"
	"time"
)

type TelnetClient interface {
	Connect() error
	io.Closer
	Send() error
	Receive() error
}

type telnetClient struct {
	address string
	timeout time.Duration
	conn    net.Conn
	in      io.Reader
	out     io.Writer
}

func NewTelnetClient(address string, timeout time.Duration, in io.Reader, out io.Writer) TelnetClient {
	return &telnetClient{
		address: address,
		timeout: timeout,
		in:      in,
		out:     out,
	}
}

func (t *telnetClient) Connect() error {
	conn, err := net.DialTimeout("tcp", t.address, t.timeout)
	if err != nil {
		return fmt.Errorf("error connecting to telnet server: %w", err)
	}
	t.conn = conn
	return nil
}

func (t *telnetClient) Send() error {
	scanner := bufio.NewScanner(t.in)
	retry := 3
	for scanner.Scan() {
		if t.conn == nil {
			return fmt.Errorf("no connection to telnet server")
		}
		for attempt := 1; attempt <= retry; attempt++ {
			_, err := t.conn.Write(append(scanner.Bytes(), '\n'))
			if err != nil {
				fmt.Printf("error writing to telnet server (attempt %d of %d): %s", attempt, retry, err)
				if attempt == retry {
					return fmt.Errorf("failed to send msg to server (attempt %d of %d)", attempt, retry)
				}
			} else {
				break
			}
		}
	}
	if errors.Is(scanner.Err(), io.EOF) {
		fmt.Println("telnet client disconnected")
	}
	return scanner.Err()
}

func (t *telnetClient) Receive() error {
	reader := bufio.NewReader(t.conn)
	for {
		line, err := reader.ReadBytes('\n')
		if len(line) != 0 {
			_, _ = t.out.Write(line)
		}
		if err != nil {
			if err == io.EOF {
				return nil
			}
			return fmt.Errorf("error reading from telnet server: %w", err)
		}
	}
}

func (t *telnetClient) Close() error {
	if t.conn != nil {
		_ = t.conn.Close()
	}
	return nil
}
