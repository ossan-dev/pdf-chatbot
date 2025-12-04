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
		Model:  "qwen3-vl:2b",
		Prompt: "Extract from the image the name of Ivan's favorite football team.",
		// Prompt: "Estrai dall'immagine il nome del team di calcio preferito di Ivan Pesenti.",
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

	// get response from Ollama
	err = client.Generate(ctx, req, respFunc)
	if err != nil {
		panic(err)
	}

}
