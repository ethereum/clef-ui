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
	readyToClose := make(chan bool)

	// notify appCancel channel on Ctrl + C
	signal.Notify(appCancel, os.Interrupt)

	// trap Ctrl+C and call cancel on the context
	ctx, cancel := context.WithCancel(context.Background())

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

	// Watch for os interrupt
	go func() {
		select {
		case <-appCancel:
			cancel()
			signal.Stop(appCancel)
			readyToClose <- true
		case <-uiClose:
			cancel()
			readyToClose <- true
		}
	}()

	clefUi.Start()

	// Exit when done
	<-readyToClose
	log.Println("Clef UI is terminated.")
}
