package main

import (
	"os"
	"context"
	"os/signal"
	"fmt"
	"log"
	"github.com/kyokan/clef-ui/pkg/rpc"
	"github.com/kyokan/clef-ui/pkg/clefclient"
)

func main() {
	// trap Ctrl+C and call cancel on the context
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	cancelChannel := make(chan os.Signal, 1)
	done := make(chan bool)
	signal.Notify(cancelChannel, os.Interrupt)

	stdin, stdout, stderr, err := clefclient.StartClef()
	if err != nil {
		log.Panicf("Cannot stard clef: %s", err)
		return
	}

	server := rpc.NewServer()
	server.ListenStdIO(ctx, stdin, stdout, stderr)

	// Watch for os interrupt
	go func() {
		<-cancelChannel
		cancel()
		fmt.Print("\n")
		log.Print("Stopped Clef UI.")
		signal.Stop(cancelChannel)
		done <- true
	}()

	// Exit when done
	<-done
}
