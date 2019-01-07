package rpc

import (
	"context"
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
		ctx:    ctx,
		stdin:  stdin,
		stdout: stdout,
	}

	go func() {
		<-s.ctx.Done()
		log.Println("Stopped RPC Request Handling.")
		s.stopHandling = true
	}()

	return s
}

// meddlingCodec is used since the native json-rpc server expects the request
// methods to be on the type <servicename>.<methodname> , separated with a dot (.).
// This codec artificially adds a servicename
type meddlingCodec struct {
	backend rpc.ServerCodec
}

func (c *meddlingCodec) ReadRequestHeader(request *rpc.Request) error {
	err := c.backend.ReadRequestHeader(request)
	if err == nil {
		request.ServiceMethod = fmt.Sprintf("ClefService.%s", request.ServiceMethod)
	}
	return err
}
func (c *meddlingCodec) Close() error {
	return c.backend.Close()
}
func (c *meddlingCodec) ReadRequestBody(a interface{}) error {
	return c.backend.ReadRequestBody(a)
}
func (c *meddlingCodec) WriteResponse(r *rpc.Response, v interface{}) error {
	return c.backend.WriteResponse(r, v)
}

func (s *Server) Start(clefUI ui.ClefUI) {
	rpc.Register(&ClefService{ui: clefUI})
	codec := &meddlingCodec{
		backend: jsonrpc2.NewServerCodecContext(s.ctx, NewRWCloseCombiner(s.stdin, s.stdout), nil),
	}
	go rpc.ServeCodec(codec)
}
