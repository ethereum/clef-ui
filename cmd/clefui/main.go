package main

import (
	"os"
 //  "io"
  // "bytes"
  // "encoding/json"
  "context"
  "os/signal"
  "fmt"
	"log"
	// "os/exec"
  // "bufio"
  "github.com/kyokan/clef-signer/pkg/rpc"
  // "net"
  // "net/rpc"
  // "net/rpc/jsonrpc"
  // "sync"
  // "strings"
	// "github.com/therecipe/qt/core"
	// "github.com/therecipe/qt/widgets"
)

func main() {
  // trap Ctrl+C and call cancel on the context
  ctx := context.Background()
  ctx, cancel := context.WithCancel(ctx)
  cancelChannel := make(chan os.Signal, 1)
  done := make(chan bool)
  signal.Notify(cancelChannel, os.Interrupt)

  stdin, stdout, stderr := rpc.StartClef()

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
