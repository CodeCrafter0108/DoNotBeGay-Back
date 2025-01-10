package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/eriscrypt/donotbegay/randomizer/internal"
	"github.com/robfig/cron/v3"
)

func main() {
	// Create scheduler with seconds precision
	c := cron.New(cron.WithSeconds())

	// Add task to scheduler
	_, err := c.AddFunc("@every 1h", internal.RandomCall)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Start scheduler
	c.Start()
	log.Println("scheduler started")

	// Create context for graceful shutdown
	ctx, cancel := context.WithCancel(context.Background())

	// Create channel for OS signals
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt, syscall.SIGTERM)

	// Listen for OS signals in separate goroutine
	go func() {
		<-signalChan
		fmt.Println("received shutdown signal")
		cancel()
	}()

	// Wait for context to be done
	<-ctx.Done()

	// Stop scheduler
	fmt.Println("stopping scheduler...")

	_, stopCancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer stopCancel()
	c.Stop()

	fmt.Println("scheduler stopped gracefully")
}
