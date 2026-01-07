package main

import (
	"context"
	"fmt"

	"github.com/ollama/ollama/api"
)

// TODO: emulate a chat here by questioning and answering with the AI-support.

func main() {
	ctx, cancelFunc := context.WithCancel(context.Background())
	defer cancelFunc()

	// get Ollama client
	client, err := api.ClientFromEnvironment()
	if err != nil {
		panic(err)
	}

	if err := client.Heartbeat(ctx); err != nil {
		panic(err)
	}

	// prepare Ollama request
	falsePtr := false
	req := &api.ChatRequest{
		Model: "qwen3-vl:235b-cloud",
		Messages: []api.Message{
			{Role: "system",
				Content: `You're Rob Pike. You're talking to a junior software engineer who cheers for Java.
Use anedoctes, jokes, be creative, but concise.`,
			},
			{
				Role:    "user",
				Content: "Why Go is coolest programming language?",
			},
		},
		Stream: &falsePtr,
	}

	// set handler for the response
	chatFunc := func(resp api.ChatResponse) error {
		fmt.Println(resp.Message.Content)
		return nil
	}

	// get response from Ollama
	err = client.Chat(ctx, req, chatFunc)
	if err != nil {
		panic(err)
	}

}
