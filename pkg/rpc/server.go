import (
  "encoding/json"
)

type Rpc struct {
  Jsonrpc   string
  Method    string
  Id        *json.RawMessage
  Params    *json.RawMessage
}

func ListenStdIO() {
  // codec := json.NewCodec()

  // Create a new codec request.
  // codecReq := codec.NewRequest(r)

  req := new(serverRequest)
  err := json.NewDecoder(r.Body).Decode(req)
  r.Body.Close()
  return &CodecRequest{request: req, err: err}



  // Get service method to be called.
  method, errMethod := codecReq.Method()
  if errMethod != nil {
    codecReq.WriteError(w, 400, errMethod)
    return
  }
  serviceSpec, methodSpec, errGet := s.services.get(method)
  if errGet != nil {
    codecReq.WriteError(w, 400, errGet)
    return
  }
  // Decode the args.
  args := reflect.New(methodSpec.argsType)
  if errRead := codecReq.ReadRequest(args.Interface()); errRead != nil {
    codecReq.WriteError(w, 400, errRead)
    return
  }

  // Call the registered Intercept Function
  if s.interceptFunc != nil {
    req := s.interceptFunc(&RequestInfo{
      Request: r,
      Method:  method,
    })
    if req != nil {
      r = req
    }
  }
  // Call the registered Before Function
  if s.beforeFunc != nil {
    s.beforeFunc(&RequestInfo{
      Request: r,
      Method:  method,
    })
  }

  // Call the service method.
  reply := reflect.New(methodSpec.replyType)
  errValue := methodSpec.method.Func.Call([]reflect.Value{
    serviceSpec.rcvr,
    reflect.ValueOf(r),
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
  w.Header().Set("x-content-type-options", "nosniff")
  // Encode the response.
  if errResult == nil {
    codecReq.WriteResponse(w, reply.Interface())
  } else {
    codecReq.WriteError(w, 400, errResult)
  }

  // Call the registered After Function
  if s.afterFunc != nil {
    s.afterFunc(&RequestInfo{
      Request:    r,
      Method:     method,
      Error:      errResult,
      StatusCode: 200,
    })
  }
}