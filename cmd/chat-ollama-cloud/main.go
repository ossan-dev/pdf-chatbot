package main

import (
	"context"
	"fmt"

	"github.com/ollama/ollama/api"
)

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

	falsePtr := false
	req := &api.ChatRequest{
		Model: "qwen3-vl:235b-cloud",
		Messages: []api.Message{
			{Role: "user", Content: "Why Go is coolest programming language?"},
		},
		Stream: &falsePtr,
	}

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
