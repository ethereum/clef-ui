type ClefRpc struct {

}

type OnSignerStartupArgs struct {

}

type OnSignerStartupReply struct {

}

func (c* ClefRpc) OnSignerStartup(args* OnSignerStartupArgs, reply* OnSignerStartupReply) {
  fmt.Println("OnSignerStartup getting called")
}
