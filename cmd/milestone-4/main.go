package main

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/ollama/ollama/api"
)

func main() {
	ctx, cancelFunc := context.WithCancel(context.Background())
	defer cancelFunc()

	// 1. get Ollama client
	client, err := api.ClientFromEnvironment()
	if err != nil {
		panic(err)
	}

	if err := client.Heartbeat(ctx); err != nil {
		panic(err)
	}

	// 2. preparing the image
	imgFile, err := os.ReadFile("/home/ossan/Projects/pdf-chatbot/imgs/booking.png")
	if err != nil {
		panic(err)
	}

	// setup Ollama request
	messages := []api.Message{
		{
			Role:    "system",
			Content: "You're an assistant. Extract information from the image.",
		},
		{Role: "user",
			Content: "Extract how much did they pay for the accommodation.",
			Images:  []api.ImageData{imgFile},
		},
	}

	falsePtr := false
	req := &api.ChatRequest{
		Model:    "qwen3-vl:2b",
		Messages: messages,
		Stream:   &falsePtr,
	}

	// 4. setup Ollama response
	chatFunc := func(resp api.ChatResponse) error {
		fmt.Print(resp.Message.Content)
		if resp.Done {
			req.Messages = append(req.Messages, resp.Message)
		}
		return nil
	}

	// kick-off conversation
	err = client.Chat(ctx, req, chatFunc)
	if err != nil {
		if err != io.EOF {
			panic(err)
		}
	}

	// chatting with Ollama
	scanner := bufio.NewScanner(os.Stdin)
	for {
		// get user's prompt
		fmt.Print("\nUser > ")
		if !scanner.Scan() {
			break
		}
		prompt := scanner.Text()
		if strings.TrimSpace(prompt) == "" {
			continue
		}

		// append user prompt to chat history
		req.Messages = append(req.Messages, api.Message{
			Role:    "user",
			Content: prompt,
		})

		fmt.Print("Assistant >")
		err = client.Chat(ctx, req, chatFunc)
		if err != nil {
			if err != io.EOF {
				panic(err)
			}
		}
	}
}
