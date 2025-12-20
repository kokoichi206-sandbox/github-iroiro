package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	fmt.Fprintf(os.Stderr, "PID: %d\n", os.Getpid())
	fmt.Fprintf(os.Stderr, "PPID: %d\n", os.Getppid())
	fmt.Fprintf(os.Stderr, "PGID: %d\n", syscall.Getpgrp())
	fmt.Fprintln(os.Stderr, "Waiting for signals...")

	sigCh := make(chan os.Signal, 1)

	// すべてのシグナルを受け取る
	signal.Notify(sigCh)

	// 定期的に生存確認を出力
	go func() {
		ticker := time.NewTicker(1 * time.Second)
		defer ticker.Stop()
		for t := range ticker.C {
			fmt.Fprintf(os.Stderr, "Still alive at %s (PID=%d, PPID=%d)\n", t.Format("15:04:05"), os.Getpid(), os.Getppid())
		}
	}()

	for {
		sig := <-sigCh
		fmt.Fprintf(os.Stderr, "Received signal: %v at %s\n", sig, time.Now().Format("2006-01-02 15:04:05"))
	}
}
