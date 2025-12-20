package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Printf("os.Stdin: %+v\n", os.Stdin)
	fmt.Printf("os.Stdin.Fd(): %v\n", os.Stdin.Fd())
	fmt.Printf("os.Stdin.Name(): %v\n", os.Stdin.Name())

	// 切り替え: "bufio" または "goroutine"
	mode := "goroutine"
	if len(os.Args) > 1 {
		mode = os.Args[1]
	}

	fmt.Printf("Mode: %s\n", mode)

	switch mode {
	case "bufio":
		modeBufio()
	case "goroutine":
		modeGoroutine()
	default:
		fmt.Printf("Unknown mode: %s\n", mode)
		fmt.Println("Usage: main [bufio|goroutine]")
	}
}

// modeBufio は bufio.Reader を使った対話モード
func modeBufio() {
	buf := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("> ")
		line, err := buf.ReadString('\n')
		if err != nil {
			fmt.Printf("err: %v\n", err)
			break
		}

		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}
		if line == "exit" {
			break
		}

		fmt.Printf("Received: %s\n", line)
	}
}

type readResult struct {
	line string
	err  error
}

// modeGoroutine は goroutine を使って読み込む対話モード
func modeGoroutine() {
	inputCh := make(chan readResult)

	// 別 goroutine で stdin を読み込む
	go func() {
		buf := bufio.NewReader(os.Stdin)
		for {
			line, err := buf.ReadString('\n')
			inputCh <- readResult{line: line, err: err}
			if err != nil {
				return
			}
		}
	}()

	for {
		fmt.Print("> ")

		result := <-inputCh
		if result.err != nil {
			fmt.Printf("err: %v\n", result.err)
			break
		}

		line := strings.TrimSpace(result.line)
		if line == "" {
			continue
		}
		if line == "exit" {
			break
		}

		fmt.Printf("Received: %s\n", line)
	}
}
