package rpc

import (
	"context"
	"github.com/kyokan/clef-ui/internal/ui"
	"github.com/powerman/rpc-codec/jsonrpc2"
	"io"
	"log"
	"os"

	"bytes"
	"encoding/json"
	"fmt"
	"net/rpc"
	"sync"
)

type Server struct {}

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
	if rwc.Buf.Len() == 0 {
		return 0, nil
	} else {
		log.Println(rwc.Buf.String())
	}
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

func (s *Server) handleRpc(ctx context.Context, stdin io.Writer, stdout io.Reader, clefUi ui.ClefUI) {
	done := false

	go func() {
		<-ctx.Done()
		log.Print("Stopped incoming rpc request handlers.")
		done = true
	}()

	rpc.Register(&ClefService{
		ui: clefUi,
	})
	rwc := NewRWCloseCombiner(stdin)
	enc := json.NewEncoder(rwc.Buf)
	dec := json.NewDecoder(stdout)

	go func() {
		for !done {
			var v map[string]interface{}
			if err := dec.Decode(&v); err != nil {
				log.Println("Failed to decode JSON:", err)
				continue
			}
			method, ok := v["method"]
			if !ok {
				log.Printf("Need Method.")
				continue
			}
			v["method"] = fmt.Sprintf("%s.%s", "ClefService", method)
			log.Println(method)
			if err := enc.Encode(&v); err != nil {
				log.Println("Failed to encode JSON:", err)
				continue
			}
		}
	}()

	log.Println("ready channel receive")
	go jsonrpc2.ServeConnContext(ctx, rwc)
}

func (s *Server) ListenStdIO(ctx context.Context, stdin io.WriteCloser, stdout io.ReadCloser, stderr io.ReadCloser, clefUi *ui.ClefUI) {
	// Non-blockingly echo command output to terminal
	go io.Copy(os.Stderr, stderr)
	go s.handleRpc(ctx, stdin, stdout, *clefUi)
}
