package clefclient

import (
	"os"
	"os/exec"
	"path"
	"io"
	"context"
	"log"
)

func StartClef(ctx context.Context) (io.WriteCloser, io.ReadCloser, io.ReadCloser, error) {
	goPath := os.Getenv("GOPATH")
	cmd := exec.Command(
		path.Join(goPath, "bin", "clef"),
		"--rpc",
		"--4bytedb",
		path.Join(goPath, "src", "github.com", "ethereum", "go-ethereum", "cmd", "clef", "4byte.json"),
		"--stdio-ui",
		"--ipcdisable",
	)

	// Create stdout, stderr streams of type io.Reader
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return nil, nil, nil, err
	}
	stderr, err := cmd.StderrPipe()
	if err != nil {
		return nil, nil, nil, err
	}
	stdin, err := cmd.StdinPipe()
	if err != nil {
		return nil, nil, nil, err
	}

	// Start command
	err = cmd.Start()
	if err != nil {
		return nil, nil, nil, err
	}

	go func() {
		<-ctx.Done()
		log.Println("Terminating clef.")
		err := cmd.Process.Signal(os.Interrupt)
		if err != nil {
			log.Println("Failed to terminate Clef cleanly:", err)
		}
	}()

	return stdin, stdout, stderr, nil
}
