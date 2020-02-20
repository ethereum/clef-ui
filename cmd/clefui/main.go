package main

import (
	"context"
	"log"
	"io"
	"os"
	"os/signal"

	"github.com/ethereum/clef-ui/external"
	"github.com/ethereum/clef-ui/internal/ui"

	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"github.com/therecipe/qt/qml"
)

func main() {
	// Make all the channels
	appCancel := make(chan os.Signal, 1)
	// readyToStart := make(chan string)

	// notify appCancel channel on Ctrl + C
	signal.Notify(appCancel, os.Interrupt)

	// trap Ctrl+C and call cancel on the context
	ctx, cancel := context.WithCancel(context.Background())

	// go func() {
	// 	select {
	// 	case gopath := <-readyToStart:
	// 		log.Println(gopath)
	// 		// Start Clef Client
	// 		stdin, stdout, stderr, err := external.StartClef(ctx, gopath)
	// 		if err != nil {
	// 			log.Panicf("Cannot start clef: %s", err)
	// 			return
	// 		}

	// 		// Just Copy stderr to terminal for now
	// 		// TODO: Handle Standard Error properly
	// 		go io.Copy(os.Stderr, stderr)

	// 		server := external.NewServer(ctx, stdin, stdout)
	// 		server.Start(*clefUi)
	// 	}
	// }()

	// Watch for os interrupt
	go func() {
		select {
		case <-appCancel:
			cancel()
			signal.Stop(appCancel)
			exit(0)
		}
	}()

	startUi(ctx)
}

func startUi(ctx context.Context) {
	core.QCoreApplication_SetAttribute(core.Qt__AA_EnableHighDpiScaling, true)
	app := gui.NewQGuiApplication(len(os.Args), os.Args)
	app.ConnectLastWindowClosed(func() {
		exit(0)
	})
	
	engine := qml.NewQQmlApplicationEngine(nil)

	login := ui.NewLogin(engine.RootContext(), "LoginContext")
	// engine.Load(core.NewQUrl3("qrc:/qml/main.qml", 0))
	engine.Load(core.NewQUrl3("internal/ui/qml/main.qml", 0))

	go func() {
		select {
		case gopath := <-login.InitiateServerRequest:
			log.Println("signal working")
			stdin, stdout, stderr, err := external.StartClef(ctx, gopath)
			if err != nil {
				log.Panicf("Cannot start clef: %s", err)
				return
			}

			// Just Copy stderr to terminal for now
			// TODO: Handle Standard Error properly
			go io.Copy(os.Stderr, stderr)

			server := external.NewServer(ctx, stdin, stdout)
			server.Start()
			log.Println("Server started")
		}
	}()

	gui.QGuiApplication_Exec()
}

func exit(returnCode int) {
	log.Println("Clef UI is terminated.")
	core.QCoreApplication_Exit(returnCode)
}
