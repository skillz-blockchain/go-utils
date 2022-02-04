package eth2http

import (
	"context"
	"net/http"

	"github.com/Azure/go-autorest/autorest"

	httppreparer "github.com/skillz-blockchain/go-utils/http/preparer"
)

// Client provides methods to connect to an Ethereum 2.0 Beacon chain node
type Client struct {
	client autorest.Sender
}

// NewClient creates a new client
func NewClient(client autorest.Sender) *Client {
	return &Client{
		client: client,
	}
}

// NewClient creates a client connecting to an Ethereum 2.0 Beacon chain node at given addr
func NewClientFromAddress(addr string) *Client {
	return NewClient(autorest.Client{
		Sender:           http.DefaultClient,
		RequestInspector: httppreparer.WithBaseURL(addr),
	})
}

func newRequest(ctx context.Context) *http.Request {
	req, _ := http.NewRequestWithContext(ctx, "", "", nil)
	return req
}

func inspectResponse(resp *http.Response, msg interface{}) error {
	return autorest.Respond(
		resp,
		WithBeaconErrorUnlessOK(),
		autorest.ByUnmarshallingJSON(msg),
		autorest.ByClosing(),
	)
}