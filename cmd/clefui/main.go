package main

import (
	"context"
	"github.com/kyokan/clef-ui/internal/ui"
	"github.com/kyokan/clef-ui/pkg/clefclient"
	"github.com/kyokan/clef-ui/pkg/rpc"
	"io"
	"log"
	"os"
	"os/signal"
)

func main() {
	// Make all the channels
	uiClose := make(chan bool)
	appCancel := make(chan os.Signal, 1)
	readyToClose := make(chan bool)

	// notify appCancel channel on Ctrl + C
	signal.Notify(appCancel, os.Interrupt)

	// trap Ctrl+C and call cancel on the context
	ctx, cancel := context.WithCancel(context.Background())

	// Start Clef Client
	stdin, stdout, stderr, err := clefclient.StartClef(ctx)
	if err != nil {
		log.Panicf("Cannot start clef: %s", err)
		return
	}

	// Just Copy stderr to terminal for now
	// TODO: Handle Standard Error properly
	go io.Copy(os.Stderr, stderr)

	clefUi := ui.NewClefUI(ctx, uiClose)

	server := rpc.NewServer(ctx, stdin, stdout)
	server.Start(*clefUi)

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

	//ui.Start(ctx, stopChan)

	clefUi.App.Exec()
	// Exit when done
	<-readyToClose
	log.Println("Clef UI is terminated.")
}
