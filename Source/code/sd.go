package buttontester

import (
	"context"
	"net/http"

	"github.com/FlowingSPDG/streamdeck"
)

const (
	ActionTest = "dev.flowingspdg.buttontester.test"
)

// ButtonTester is a minimal StreamDeck client that logs button events.
type ButtonTester struct {
	sd *streamdeck.Client
	ht *http.Client
}

// NewButtonTester constructs a new plugin instance and registers handlers.
func NewButtonTester(ctx context.Context, params streamdeck.RegistrationParams) *ButtonTester {
	client := &ButtonTester{
		sd: streamdeck.NewClient(ctx, params),
		ht: http.DefaultClient,
	}

	action := client.sd.Action(ActionTest)
	action.RegisterHandler(streamdeck.WillAppear, client.WillAppear)
	action.RegisterHandler(streamdeck.WillDisappear, client.WillDisappear)
	action.RegisterHandler(streamdeck.KeyDown, client.KeyDown)
	action.RegisterHandler(streamdeck.KeyUp, client.KeyUp)

	return client
}

// Run starts the StreamDeck client event loop.
func (b *ButtonTester) Run(ctx context.Context) error {
	return b.sd.Run(ctx)
}
