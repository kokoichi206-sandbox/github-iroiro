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
	fmt.Fprintln(os.Stderr, "Waiting for signals... (Ctrl+C to send SIGINT)")
	fmt.Fprintln(os.Stderr, "Try: kill -TERM <pid>, kill -HUP <pid>, kill -USR1 <pid>, etc.")

	sigCh := make(chan os.Signal, 1)

	// 受け取りたいシグナルを登録
	signal.Notify(sigCh,
		syscall.SIGINT,  // Ctrl+C
		syscall.SIGTERM, // kill コマンドのデフォルト
		syscall.SIGHUP,  // ハングアップ
		syscall.SIGQUIT, // Ctrl+\
		syscall.SIGUSR1, // ユーザー定義シグナル 1
		syscall.SIGUSR2, // ユーザー定義シグナル 2
	)

	for {
		sig := <-sigCh
		// stderr に出力 (バッファリングされない)
		fmt.Fprintf(os.Stderr, "Received signal: %v (%d) at %s\n", sig, sig.(syscall.Signal), time.Now().Format("2006-01-02 15:04:05"))

		// SIGINT または SIGTERM で終了
		if sig == syscall.SIGINT || sig == syscall.SIGTERM {
			fmt.Fprintln(os.Stderr, "Exiting...")
			break
		}
	}
}
