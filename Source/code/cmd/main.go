package main

import (
	"context"
	"log"
	"os"

	"github.com/FlowingSPDG/streamdeck"

	buttontester "github.com/FlowingSPDG/streamdeckbuttontester/Source/code"
)

func main() {
	log.Println("Starting StreamDeck Button Tester...")

	ctx := context.Background()
	if err := run(ctx); err != nil {
		log.Fatalf("%v\n", err)
	}
}

func run(ctx context.Context) error {
	params, err := streamdeck.ParseRegistrationParams(os.Args)
	if err != nil {
		return err
	}

	client := buttontester.NewButtonTester(ctx, params)
	return client.Run(ctx)
}
