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

	buf := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("> ")
		line, err := buf.ReadString('\n')
		if err != nil { // io.EOF または他のエラー
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
