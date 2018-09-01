package rpc

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/kyokan/clef-ui/internal/ui"
	"github.com/powerman/rpc-codec/jsonrpc2"
	"io"
	"log"
	"net/rpc"
)

type Server struct {
	ctx          context.Context
	stdin        io.Writer
	stdout       io.Reader
	stopHandling bool
}

func NewServer(ctx context.Context, stdin io.Writer, stdout io.Reader) *Server {
	s := &Server{
		ctx:ctx,
		stdin:stdin,
		stdout:stdout,
	}

	go func() {
		<-s.ctx.Done()
		log.Println("Stopped RPC Request Handling.")
		s.stopHandling = true
	}()

	return s
}

func (s *Server) Start(clefUI ui.ClefUI) {
	rpc.Register(&ClefService{ ui: clefUI })
	rwc := NewRWCloseCombiner(s.stdin)
	enc := json.NewEncoder(rwc.Buf)
	dec := json.NewDecoder(s.stdout)

	go func() {
		for !s.stopHandling {
			// Decode JSON
			var v map[string]interface{}
			if err := dec.Decode(&v); err != nil {
				log.Println("Failed to decode JSON:", err)
				continue
			}
			method, ok := v["method"]
			if !ok {
				log.Printf("No RPC method found in JSON.")
				continue
			}

			// Change method name from "RPCMethodName" to "ClefService.RPCMethodName"
			// The jsonrpc2 library we use expect a dot-separated method name
			v["method"] = fmt.Sprintf("%s.%s", "ClefService", method)
			log.Printf("Incoming RPC Request - %v", method)
			if err := enc.Encode(&v); err != nil {
				log.Println("Failed to encode JSON:", err)
				continue
			}
		}
	}()

	go jsonrpc2.ServeConnContext(s.ctx, rwc)
}
