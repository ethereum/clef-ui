package rpc

import (
  // "encoding/json"
  "fmt"
)

type RpcHandlers struct {

}

type OnSignerStartupInfo struct {
  Extapi_http     string
  Extapi_ipc      string
  Extapi_version  string
  Intapi_version  string
}

type OnSignerStartupParam struct {
  Info      OnSignerStartupInfo
}

type OnSignerStartupArgs struct {
  Jsonrpc   string
  Method    string
  Id        int
  Params    []OnSignerStartupParam
}

type OnSignerStartupReply struct {

}

type ApproveListingArgs struct {

}

type ApproveListingReply struct {

}

func (c* RpcHandlers) OnSignerStartup(args* OnSignerStartupArgs, reply* OnSignerStartupReply) error {
  // fmt.Println("OnSignerStartup getting called", args, reply)
  fmt.Println(args.Params[0])
  return nil
}

func (c* RpcHandlers) ApproveListing(args* ApproveListingArgs, reply* ApproveListingReply) error {
  fmt.Println("ApproveListing getting called")
  return nil
}
