package main

import (
	"context"
	"log"
	"os"
	"os/signal"

	"github.com/kyokan/clef-ui/pkg/clefclient"
	"github.com/kyokan/clef-ui/pkg/rpc"
	"github.com/kyokan/clef-ui/internal/ui"
)

func main() {
	// trap Ctrl+C and call cancel on the context
	ctx, cancel := context.WithCancel(context.Background())
	stopChan := make(chan bool)
	cancelChannel := make(chan os.Signal, 1)
	done := make(chan bool)
	signal.Notify(cancelChannel, os.Interrupt)

	stdin, stdout, stderr, err := clefclient.StartClef(ctx)
	if err != nil {
		log.Panicf("Cannot start clef: %s", err)
		return
	}

	clefUi := ui.NewClefUI(ctx, stopChan)

	server := rpc.NewServer()
	server.ListenStdIO(ctx, stdin, stdout, stderr, clefUi)

	// Watch for os interrupt
	go func() {
		select {
		case <-cancelChannel:
			cancel()
			signal.Stop(cancelChannel)
			done <- true
		case <-stopChan:
			cancel()
			done <- true
		}
	}()

	//ui.Start(ctx, stopChan)
	log.Println("Start App Exec")
	clefUi.View.Show()
	clefUi.App.Exec()
	// Exit when done
	<-done
}
