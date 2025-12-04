package main

import (
	"context"
	"flag"
	"fmt"
	"io"
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

	for {
		// getting the request
		var prompt string
		fmt.Println("Hey there, what do you want to know about the image you have uploaded?")
		n, err := fmt.Scan(&prompt)
		if n == 0 || err == io.EOF {
			return
		}
		if err != nil {
			// panic(err)
			fmt.Fprintf(os.Stderr, "failing to read the prompt with error: %v\n", err.Error())
			return
		}

		// prepare Ollama request
		req := &api.GenerateRequest{
			Model:  "qwen3-vl:2b",
			Prompt: prompt,
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

}
