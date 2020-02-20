package external

import (
	"context"
	// "encoding/json"
	"fmt"
	"io"
	"log"
	// "math/big"
	// "strings"

	// todo: gfz move out all calls to UI
	// "github.com/ethereum/clef-ui/internal/ui"
	// "github.com/ethereum/clef-ui/internal/utils"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/ethereum/go-ethereum/signer/core"
)

type Server struct {
	client *rpc.Client
	api ClefUIAPI
}

// NewServer starts up a listener for RPC messages from the given channels
func NewServer(ctx context.Context, remoteIn io.Writer, remoteOut io.Reader) *Server {
	c, err := rpc.DialIO(ctx, remoteOut, remoteIn)
	if err != nil {
		log.Fatalf("Could not create stdio client", "err", err)
	}
	return &Server{
		client: c,
	}
}

// todo: gfz move out all calls to UI
func (s *Server) Start() {
	s.api = ClefUIAPI{ test: "test" }
	s.client.RegisterName("ui", s.api)
}

// clefUIAPI implemnents the handlers for json-rpc invocations, gfz: not sure where this is defined
type ClefUIAPI struct {
	test string
	// todo: gfz move out all calls to UI
	// ui ui.ClefUI
}

func (c *ClefUIAPI) OnSignerStartup(info core.StartupInfo) {
	fmt.Printf("clefService.OnSignerStartup sending signal \n")
	// todo: gfz move out all calls to UI
// 	// c.ui.BackToMain <- true
}

func (c *ClefUIAPI) ShowError(message core.Message) {
	// text := strings.Replace(message.Text, "\u003c", "less than ", -1)
	// todo: gfz move out all calls to UI
	// c.ui.ErrorDialog <- text
}

func (c *ClefUIAPI) ShowInfo(message core.Message) {
	// TODO! Separate info and error
	// text := strings.Replace(message.Text, "\u003c", "less than ", -1)
	// todo: gfz move out all calls to UI
	// c.ui.ErrorDialog <- text
}

func (c *ClefUIAPI) OnUserInputrequest(message core.UserInputRequest) (*core.UserInputResponse, error) {
	// response, err := c.ui.RequestUserInput(message.Title, message.Prompt, message.IsPassword)
	// if err != nil {
	// 	return nil, err
	// }
	return &core.UserInputResponse{
		// Text: response,
	}, nil
}

func (c *ClefUIAPI) ApproveNewAccount(p core.NewAccountRequest) (response *core.NewAccountResponse, err error) {
	ch := make(chan *core.NewAccountResponse)
	// item := ui.IncomingRequestItem{
	// 	From:        " - ",
	// 	Description: "Request for new account creation",
	// 	RPC: &ui.ApproveNewAccountRequest{
	// 		Params:     &p,
	// 		ResponseCh: ch,
	// 	},
	// }

	// // todo: gfz move out all calls to UI
	// c.ui.IncomingRequest <- item
	r := <-ch
	// item.Remove()
	return r, nil
}

func (c *ClefUIAPI) ApproveTx(p core.SignTxRequest) (*core.SignTxResponse, error) {
	// val := big.Int(p.Transaction.Value)
	// desc := fmt.Sprintf("Transaction: %s to 0x%x... ",
	// 	clefutils.DefaultFormat(&val),
	// 	p.Transaction.To.Address().Bytes()[:4])

	ch := make(chan *core.SignTxResponse)
	// item := ui.IncomingRequestItem{
	// 	From:        p.Transaction.From.Original(),
	// 	Description: desc,
	// 	RPC: &ui.ApproveTxRequest{
	// 		Params:     &p,
	// 		ResponseCh: ch,
	// 	},
	// }
	// // todo: gfz move out all calls to UI
	// // Send to UI
	// c.ui.IncomingRequest <- item
	// // Wait for response
	response := <-ch
	// // Remove from UI
	// item.Remove()

	// // debug print TODO remove
	// d, e := json.Marshal(&response)
	// fmt.Printf("<-- %s (%v)\n", d, e)
	// //
	return response, nil
}

func (c *ClefUIAPI) ApproveListing(p core.ListRequest) (*core.ListResponse, error) {
	ch := make(chan *core.ListResponse)
	// desc := "Request to list accounts"
	// if len(p.Meta.Origin) > 0 {
	// 	desc = fmt.Sprintf("Request to list accounts (orgin %s)", p.Meta.Origin)
	// } else if len(p.Meta.UserAgent) > 0 {
	// 	desc = fmt.Sprintf("Request to list accounts (ua: %s)", p.Meta.UserAgent)
	// }
	// item := ui.IncomingRequestItem{
	// 	From:        " - ",
	// 	Description: desc,
	// 	RPC: &ui.ApproveListingRequest{
	// 		Params:     p,
	// 		ResponseCh: ch,
	// 	},
	// }

	// // todo: gfz move out all calls to UI
	// c.ui.IncomingRequest <- item
	response := <-ch
	// item.Remove()
	// // debug print TODO remove
	// d, e := json.Marshal(&response)
	// fmt.Printf("<-- %s (%v)\n", d, e)
	// //
	return response, nil
}

func (c *ClefUIAPI) ApproveSignData(p *core.SignDataRequest) (*core.SignDataResponse, error) {
	return nil, fmt.Errorf("Not implemented")
}
