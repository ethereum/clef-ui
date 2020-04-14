package main

import (
	"context"
	"io"
	"log"
	"os"
	"os/signal"

	"github.com/ethereum/clef-ui/external"
	"github.com/ethereum/clef-ui/internal/ui"
)

func main() {
	// Make all the channels
	uiClose := make(chan bool)
	appCancel := make(chan os.Signal, 1)
	readyToStart := make(chan string)

	// notify appCancel channel on Ctrl + C
	signal.Notify(appCancel, os.Interrupt)

	// trap Ctrl+C and call cancel on the context
	ctx, cancel := context.WithCancel(context.Background())
	_ = cancel

	clefUi := ui.NewClefUI(ctx, uiClose, readyToStart)

	go func() {
		select {
		case gopath := <-readyToStart:
			log.Println("ready to start")
			log.Println(gopath)
			// Start Clef Client
			stdin, stdout, stderr, err := external.StartClef(ctx, gopath)
			if err != nil {
				log.Panicf("Cannot start clef: %s", err)
				return
			}

			// Just Copy stderr to terminal for now
			// TODO: Handle Standard Error properly
			go io.Copy(os.Stderr, stderr)

			server := external.NewServer(ctx, stdin, stdout)
			server.Start(*clefUi)
		}
	}()

	clefUi.Start()

	log.Println("Clef UI is terminated.")
}
