package clefclient

import (
	"os"
	"os/exec"
	"path"
	"io"
	"log"
	"github.com/therecipe/qt/core"
)

func StartClef() (io.WriteCloser, io.ReadCloser, io.ReadCloser, error) {
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

	p := &core.QProcess{}
	log.Print(p)

	return stdin, stdout, stderr, nil
}