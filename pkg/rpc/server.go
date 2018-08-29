package rpc

import (
  "bufio"
  "context"
  "encoding/json"
  // "fmt"
  "log"
  "io"
  "reflect"
  "os/exec"
  "os"
  // "net/rpc"
  // "net/rpc/jsonrpc"
)

const ClefServiceName = "clef"

type RawParamsRpc struct {
  Jsonrpc   string
  Method    string
  Id        *json.RawMessage
  Params    *json.RawMessage
  Raw       []byte
}

type Server struct {
  serviceMap *serviceMap
}

func NewServer() *Server {
  sMap := &serviceMap{}
  if err := sMap.register(&RpcHandlers{}, ClefServiceName); err != nil {
    panic(err)
  }

  return &Server{
    serviceMap: sMap,
  }
}

func checkError(err error) {
  if err != nil {
    log.Fatalf("Error: %s", err)
  }
}

func StartClef() (io.WriteCloser, io.ReadCloser, io.ReadCloser) {
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
  stdin, err := cmd.StdinPipe()
  checkError(err)

  // Start command
  err = cmd.Start()
  checkError(err)

  return stdin, stdout, stderr
}

func (s *Server) scanStdout(ctx context.Context, incomingRpcChannel chan RawParamsRpc, stdout io.ReadCloser) {
  // var allText []string
  var done = false
  // var j Rpc

  buff := bufio.NewScanner(stdout)

  go func() {
    <-ctx.Done()
    log.Print("Stopped scanning Standard Output for new rpc request.")
    done = true
  }()

  // Iterate over buff and append content to the slice
  for !done && buff.Scan() {
      j := RawParamsRpc{}
      j.Raw = []byte(buff.Text())
      err := json.Unmarshal([]byte(buff.Text()), &j)
      checkError(err)
      log.Printf("Incoming RPC Request: %d - %+v\n", j.Id, j.Method)
      
      // allText = append(allText, buff.Text()+"\n")
      incomingRpcChannel <- j
  }
}

func (s *Server) handleRpc(ctx context.Context, incomingRpcChannel chan RawParamsRpc) {
  done := false

  go func() {
    <-ctx.Done()
    log.Print("Stopped incoming rpc request handlers.")
    done = true
  }()

  for !done {
    rpc := <-incomingRpcChannel
    serviceSpec, methodSpec, err := s.serviceMap.get(ClefServiceName, rpc.Method)
    if err != nil {
      panic(err)
    }

    args := reflect.New(methodSpec.argsType)
    err = json.Unmarshal(rpc.Raw, args.Interface())
    if err != nil {
      panic(err)
    }
    reply := reflect.New(methodSpec.replyType)
  errValue := methodSpec.method.Func.Call([]reflect.Value{
    serviceSpec.rcvr,
    args,
    reply,
  })
  // Cast the result to error if needed.
  var errResult error
  errInter := errValue[0].Interface()
  if errInter != nil {
    errResult = errInter.(error)
  }
  // Prevents Internet Explorer from MIME-sniffing a response away
  // from the declared content-type
  // Encode the response.
  if errResult == nil {
    log.Println(reply)
  } else {
    log.Println(errResult)
  }
  }
}

func (s *Server) ListenStdIO(ctx context.Context, stdin io.WriteCloser, stdout io.ReadCloser, stderr io.ReadCloser) {
  incomingRpcChannel := make(chan RawParamsRpc)

  // Non-blockingly echo command output to terminal
  go io.Copy(os.Stderr, stderr)
  go s.scanStdout(ctx, incomingRpcChannel, stdout)
  go s.handleRpc(ctx, incomingRpcChannel)
  // rpc := new(Rpc)
  // err := json.NewDecoder(r.Body).Decode(rpc)



  // Get service method to be called.
  // method, errMethod := codecReq.Method()
  // if errMethod != nil {
  //   codecReq.WriteError(w, 400, errMethod)
  //   return
  // }
  // serviceSpec, methodSpec, errGet := s.services.get(method)
  // if errGet != nil {
  //   codecReq.WriteError(w, 400, errGet)
  //   return
  // }
  // // Decode the args.
  // args := reflect.New(methodSpec.argsType)
  // if errRead := codecReq.ReadRequest(args.Interface()); errRead != nil {
  //   codecReq.WriteError(w, 400, errRead)
  //   return
  // }

  // // Call the registered Intercept Function
  // if s.interceptFunc != nil {
  //   req := s.interceptFunc(&RequestInfo{
  //     Request: r,
  //     Method:  method,
  //   })
  //   if req != nil {
  //     r = req
  //   }
  // }
  // // Call the registered Before Function
  // if s.beforeFunc != nil {
  //   s.beforeFunc(&RequestInfo{
  //     Request: r,
  //     Method:  method,
  //   })
  // }

  // // Call the service method.
  // reply := reflect.New(methodSpec.replyType)
  // errValue := methodSpec.method.Func.Call([]reflect.Value{
  //   serviceSpec.rcvr,
  //   reflect.ValueOf(r),
  //   args,
  //   reply,
  // })
  // // Cast the result to error if needed.
  // var errResult error
  // errInter := errValue[0].Interface()
  // if errInter != nil {
  //   errResult = errInter.(error)
  // }
  // // Prevents Internet Explorer from MIME-sniffing a response away
  // // from the declared content-type
  // w.Header().Set("x-content-type-options", "nosniff")
  // // Encode the response.
  // if errResult == nil {
  //   codecReq.WriteResponse(w, reply.Interface())
  // } else {
  //   codecReq.WriteError(w, 400, errResult)
  // }

  // // Call the registered After Function
  // if s.afterFunc != nil {
  //   s.afterFunc(&RequestInfo{
  //     Request:    r,
  //     Method:     method,
  //     Error:      errResult,
  //     StatusCode: 200,
  //   })
  // }
}