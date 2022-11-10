package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/shivtej1505/gosql/engine"
	"github.com/shivtej1505/gosql/engine/cli"
)

func main() {
	cleanupChan := make(chan os.Signal)
	signal.Notify(cleanupChan, syscall.SIGTERM, syscall.SIGINT)

	go func() {
		<-cleanupChan
		fmt.Println("\nbye")
		os.Exit(0)
	}()

	engine := engine.NewEngine()

	prompt := cli.NewPrompt(engine)
	prompt.StartPrompt()
}
