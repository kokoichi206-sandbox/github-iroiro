package main

import (
	"fmt"
	"os"
	"os/signal"
	"time"
)

func main() {
	fmt.Fprintf(os.Stderr, "PID: %d\n", os.Getpid())
	fmt.Fprintln(os.Stderr, "Waiting for signals... (Ctrl+C to send SIGINT)")
	fmt.Fprintln(os.Stderr, "Try: kill -TERM <pid>, kill -HUP <pid>, kill -USR1 <pid>, etc.")

	sigCh := make(chan os.Signal, 1)

	// すべてのシグナルを受け取る
	signal.Notify(sigCh)

	// 定期的に生存確認を出力
	go func() {
		ticker := time.NewTicker(1 * time.Second)
		defer ticker.Stop()
		for t := range ticker.C {
			fmt.Fprintf(os.Stderr, "Still alive at %s\n", t.Format("15:04:05"))
		}
	}()

	for {
		sig := <-sigCh
		// stderr に出力 (バッファリングされない)
		fmt.Fprintf(os.Stderr, "Received signal: %v at %s\n", sig, time.Now().Format("2006-01-02 15:04:05"))
	}
}
