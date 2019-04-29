package external

import (
	"context"
	"io"
	"log"
	"os"
	"os/exec"
	"path"
)

func StartClef(ctx context.Context, clefBin string) (io.WriteCloser, io.ReadCloser, io.ReadCloser, error) {
	cmd := exec.Command(
		clefBin,
		"--rpc",
		"--stdio-ui", "--stdio-ui-test", "--advanced",
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
