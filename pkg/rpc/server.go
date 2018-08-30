package rpc

import (
	"context"
	"log"
	"io"
	"os"

	"net/rpc"
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/powerman/rpc-codec/jsonrpc2"
	"sync"
)

type Server struct {
}

func NewServer() *Server {
	return &Server{}
}

type RWCloseCombiner struct {
	mtx         sync.Mutex
	w           io.Writer
	Buf         *bytes.Buffer
	newDataChan chan bool
}

func NewRWCloseCombiner(w io.Writer) *RWCloseCombiner {
	var buf bytes.Buffer

	return &RWCloseCombiner{
		w:   w,
		Buf: &buf,
	}
}

func (rwc *RWCloseCombiner) Read(p []byte) (int, error) {
	rwc.mtx.Lock()
	defer rwc.mtx.Unlock()
	return rwc.Buf.Read(p)
}

func (rwc *RWCloseCombiner) Write(p []byte) (int, error) {
	rwc.mtx.Lock()
	defer rwc.mtx.Unlock()
	return rwc.w.Write(p)
}

func (rwc *RWCloseCombiner) Close() error {
	return nil
}

func (s *Server) handleRpc(ctx context.Context, incomingRpcChannel chan []byte, stdin io.Writer, stdout io.Reader) {
	done := false

	go func() {
		<-ctx.Done()
		log.Print("Stopped incoming rpc request handlers.")
		done = true
	}()

	rpc.Register(&ClefService{})
	rwc := NewRWCloseCombiner(stdin)
	enc := json.NewEncoder(rwc.Buf)
	dec := json.NewDecoder(stdout)
	ready := make(chan bool)

	go func() {
		for !done {
			var v map[string]interface{}
			if err := dec.Decode(&v); err != nil {
				log.Println("Failed to decode JSON:", err)
				continue
			}
			method, ok := v["method"]
			if !ok {
				panic("need method")
			}
			v["method"] = fmt.Sprintf("%s.%s", "ClefService", method)
			if err := enc.Encode(&v); err != nil {
				log.Panic(err)
			}
			ready <- true
		}
	}()

	<-ready
	go jsonrpc2.ServeConnContext(ctx, rwc)
}

func (s *Server) ListenStdIO(ctx context.Context, stdin io.WriteCloser, stdout io.ReadCloser, stderr io.ReadCloser) {
	incomingRpcChannel := make(chan []byte)

	// Non-blockingly echo command output to terminal
	go io.Copy(os.Stderr, stderr)
	go s.handleRpc(ctx, incomingRpcChannel, stdin, stdout)
}
