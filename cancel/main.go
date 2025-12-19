package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/chzyer/readline"
)

func main() {
	fmt.Printf("os.Stdin: %+v\n", os.Stdin)
	fmt.Printf("os.Stdin.Fd(): %v\n", os.Stdin.Fd())
	fmt.Printf("os.Stdin.Name(): %v\n", os.Stdin.Name())

	// readline の設定（Terraform とほぼ同じ設定）
	l, err := readline.NewEx(&readline.Config{
		Prompt:            "> ",
		InterruptPrompt:   "^C",
		EOFPrompt:         "exit",
		HistoryFile:       "/tmp/readline.tmp", // 履歴を保存
		HistorySearchFold: true,
		Stdin:             os.Stdin,
		Stdout:            os.Stdout,
		Stderr:            os.Stderr,
	})
	if err != nil {
		panic(err)
	}
	defer l.Close()

	for {
		line, err := l.Readline()
		if err != nil { // io.EOF または readline.ErrInterrupt
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
