package main

import (
	"os"
  "io"
  // "bytes"
  "encoding/json"
  "fmt"
	"log"
	"os/exec"
  "bufio"
  // "net"
  // "net/rpc"
  // "net/rpc/jsonrpc"
  // "sync"
  // "strings"
	// "github.com/therecipe/qt/core"
	// "github.com/therecipe/qt/widgets"
)

func checkError(err error) {
    if err != nil {
        log.Fatalf("Error: %s", err)
    }
}

func startClef() {
    cmd := exec.Command(
      "/Users/chikeichan/Projects/ethereum/go-ethereum/build/bin/clef",
      "--rpc",
      "--4bytedb",
      "/Users/chikeichan/Projects/ethereum/go-ethereum/cmd/clef/4byte.json",
      "--stdio-ui",
      "--ipcdisable",
    )

    // Create stdout, stderr streams of type io.Reader
    stdout, err := cmd.StdoutPipe()
    checkError(err)
    stderr, err := cmd.StderrPipe()
    checkError(err)
    // stdin, err := cmd.StdinPipe()
    // checkError(err)

    // Start command
    err = cmd.Start()
    checkError(err)
    // Don't let main() exit before our command has finished running
    defer cmd.Wait()  // Doesn't block

    // Non-blockingly echo command output to terminal
    // go io.Copy(os.Stdout, stdout)
    go io.Copy(os.Stderr, stderr)
    // go func() {
    buff := bufio.NewScanner(stdout)
    var allText []string

    type Param struct {
      Info      interface{}
      Address   string
      Hash      string
      Message   string
      Meta      interface{}
      Raw_data  string
    }

    type Rpc struct {
      Jsonrpc   string
      Method    string
      Id        int
      Params    []Param
    }

    var j Rpc

    // Iterate over buff and append content to the slice
    for buff.Scan() {
        j = Rpc{}
        err := json.Unmarshal([]byte(buff.Text()), &j)
        checkError(err)
        
        // fmt.Printf("%+v\n", j.Params)
        fmt.Printf("%+v\n", buff.Text())
        allText = append(allText, buff.Text()+"\n")
    }

    // }()

}

func main() {
  startClef()
  fmt.Printf("Do other stuff here! No need to wait.\n\n")
	// window.Show()

	// app.Exec()
}
