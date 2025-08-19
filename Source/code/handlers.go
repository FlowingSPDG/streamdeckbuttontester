package buttontester

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/FlowingSPDG/streamdeck"
)

// WillAppear logs when the action appears on the Stream Deck.
func (b *ButtonTester) WillAppear(ctx context.Context, client *streamdeck.Client, event streamdeck.Event) error {
	msg := fmt.Sprintf("WillAppear: context=%s action=%s device=%s", event.Context, event.Action, event.Device)
	client.LogMessage(ctx, msg)
	return nil
}

// WillDisappear logs when the action disappears.
func (b *ButtonTester) WillDisappear(ctx context.Context, client *streamdeck.Client, event streamdeck.Event) error {
	msg := fmt.Sprintf("WillDisappear: context=%s action=%s device=%s", event.Context, event.Action, event.Device)
	client.LogMessage(ctx, msg)
	return nil
}

// KeyDown logs payload when a key is pressed.
func (b *ButtonTester) KeyDown(ctx context.Context, client *streamdeck.Client, event streamdeck.Event) error {
	var payload map[string]any
	_ = json.Unmarshal(event.Payload, &payload)
	msg := fmt.Sprintf("KeyDown: context=%s action=%s device=%s payload=%v", event.Context, event.Action, event.Device, payload)
	client.LogMessage(ctx, msg)
	return nil
}

// KeyUp logs payload when a key is released.
func (b *ButtonTester) KeyUp(ctx context.Context, client *streamdeck.Client, event streamdeck.Event) error {
	var payload map[string]any
	_ = json.Unmarshal(event.Payload, &payload)
	msg := fmt.Sprintf("KeyUp: context=%s action=%s device=%s payload=%v", event.Context, event.Action, event.Device, payload)
	client.LogMessage(ctx, msg)
	return nil
}
