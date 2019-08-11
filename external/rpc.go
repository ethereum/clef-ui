package external

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"math/big"
	"strings"

	"github.com/ethereum/clef-ui/internal/ui"
	"github.com/ethereum/clef-ui/internal/utils"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/ethereum/go-ethereum/signer/core"
)

type Server struct {
	client *rpc.Client
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

func (s *Server) Start(clefUI ui.ClefUI) {
	s.client.RegisterName("ui", &clefUIAPI{ui: clefUI})
}

// clefUIAPI implemnents the handlers for json-rpc invocations
type clefUIAPI struct {
	ui ui.ClefUI
}

func (c *clefUIAPI) OnSignerStartup(info core.StartupInfo) {
	fmt.Printf("clefService.OnSignerStartup sending signal \n")
	c.ui.BackToMain <- true
}

func (c *clefUIAPI) ShowError(message core.Message) {

	text := strings.Replace(message.Text, "\u003c", "less than ", -1)
	c.ui.ErrorDialog <- text
}

func (c *clefUIAPI) ShowInfo(message core.Message) {
	// TODO! Separate info and error
	text := strings.Replace(message.Text, "\u003c", "less than ", -1)
	c.ui.ErrorDialog <- text
}

func (c *clefUIAPI) OnUserInputrequest(message core.UserInputRequest) (*core.UserInputResponse, error) {
	response, err := c.ui.RequestUserInput(message.Title, message.Prompt, message.IsPassword)
	if err != nil {
		return nil, err
	}
	return &core.UserInputResponse{
		Text: response,
	}, nil

}

func (c *clefUIAPI) ApproveNewAccount(p core.NewAccountRequest) (response *core.NewAccountResponse, err error) {
	ch := make(chan *core.NewAccountResponse)
	item := ui.IncomingRequestItem{
		From:        " - ",
		Description: "Request for new account creation",
		RPC: &ui.ApproveNewAccountRequest{
			Params:     &p,
			ResponseCh: ch,
		},
	}
	c.ui.IncomingRequest <- item
	r := <-ch
	item.Remove()
	return r, nil
}

func (c *clefUIAPI) ApproveTx(p core.SignTxRequest) (*core.SignTxResponse, error) {
	val := big.Int(p.Transaction.Value)
	desc := fmt.Sprintf("Transaction: %s to 0x%x... ",
		clefutils.DefaultFormat(&val),
		p.Transaction.To.Address().Bytes()[:4])

	ch := make(chan *core.SignTxResponse)
	item := ui.IncomingRequestItem{
		From:        p.Transaction.From.Original(),
		Description: desc,
		RPC: &ui.ApproveTxRequest{
			Params:     &p,
			ResponseCh: ch,
		},
	}
	// Send to UI
	c.ui.IncomingRequest <- item
	// Wait for response
	response := <-ch
	// Remove from UI
	item.Remove()

	// debug print TODO remove
	d, e := json.Marshal(&response)
	fmt.Printf("<-- %s (%v)\n", d, e)
	//
	return response, nil
}

func (c *clefUIAPI) ApproveListing(p core.ListRequest) (*core.ListResponse, error) {
	ch := make(chan *core.ListResponse)
	desc := "Request to list accounts"
	if len(p.Meta.Origin) > 0 {
		desc = fmt.Sprintf("Request to list accounts (orgin %s)", p.Meta.Origin)
	} else if len(p.Meta.UserAgent) > 0 {
		desc = fmt.Sprintf("Request to list accounts (ua: %s)", p.Meta.UserAgent)
	}
	item := ui.IncomingRequestItem{
		From:        " - ",
		Description: desc,
		RPC: &ui.ApproveListingRequest{
			Params:     p,
			ResponseCh: ch,
		},
	}
	c.ui.IncomingRequest <- item
	response := <-ch
	item.Remove()
	// debug print TODO remove
	d, e := json.Marshal(&response)
	fmt.Printf("<-- %s (%v)\n", d, e)
	//
	return response, nil
}

func (c *clefUIAPI) ApproveSignData(p *core.SignDataRequest) (*core.SignDataResponse, error) {
	return nil, fmt.Errorf("Not implemented")
}
