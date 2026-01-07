package main

import (
	"context"
	"flag"
	"fmt"
	"os"

	"github.com/ollama/ollama/api"
)

var imagePath *string

func main() {
	ctx, cancelFunc := context.WithCancel(context.Background())
	defer cancelFunc()

	imagePath = flag.String("imagePath", "imgs/sample.png", "path of the image to load")

	flag.Parse()

	imgData, err := os.ReadFile(*imagePath)
	if err != nil {
		panic(err)
	}

	// get Ollama client
	client, err := api.ClientFromEnvironment()
	if err != nil {
		panic(err)
	}

	// prepare Ollama request
	req := &api.GenerateRequest{
		Model:  "qwen2.5vl:7b",
		Prompt: "Extract the name of the Favorite Football Team from the image.",
		Images: []api.ImageData{imgData},
	}

	// set handler for the response
	respFunc := func(resp api.GenerateResponse) error {
		if resp.Response != "" {
			fmt.Print(resp.Response)
		}
		if resp.Done {
			fmt.Println()
		}
		return nil
	}

	// get response from Ollama client
	err = client.Generate(ctx, req, respFunc)
	if err != nil {
		panic(err)
	}

}
