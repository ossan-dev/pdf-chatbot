package main

import (
	"bufio"
	"bytes"
	"context"
	"fmt"
	"image"
	"image/png"
	"io"
	"os"
	"strings"

	"github.com/ollama/ollama/api"
)

// SubImager abstracts the SubImage method on the various image type
// eg. RGBA, RGBA64, Gray, Gray16, Alpha, Alpha16, etc.
type SubImager interface {
	SubImage(image.Rectangle) image.Image
}

// GetTilesFromImg is the algorithm used to tile up the image
func GetTilesFromImg(img image.Image, tileWidth, tileHeight, overlap int) ([]image.Image, error) {
	tiles := make([]image.Image, 0, 1000)

	// 1. returns width, height
	rect := img.Bounds()
	imgWidth := rect.Dx()
	imgHeight := rect.Dy()

	// 2. moving vertically (from top to bottom)
	for y := 0; y <= imgHeight; y += tileHeight - overlap {
		// 3. moving horizontally (from left to right)
		for x := 0; x <= imgWidth; x += tileWidth - overlap {
			tileX := x
			tileY := y
			tileHeight := tileHeight
			tileWidth := tileWidth
			// 4. avoid exceeding horizontal boundary
			if tileX+tileWidth > imgWidth {
				tileWidth = imgWidth - x
			}
			// 5. avoid exceeding vertical boundary
			if tileY+tileHeight > imgHeight {
				tileHeight = imgHeight - y
			}
			// 6. defining tile
			tileBounds := image.Rect(tileX, tileY, tileX+tileWidth, tileY+tileHeight)
			// 7. slicing the current tile from the img
			tileImg := img.(SubImager).SubImage(tileBounds)
			tiles = append(tiles, tileImg)
		}
	}
	return tiles, nil
}

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
	imgFile, err := os.Open("/home/ossan/Projects/pdf-chatbot/imgs/booking.png")
	if err != nil {
		panic(err)
	}
	defer imgFile.Close()

	img, err := png.Decode(imgFile)
	if err != nil {
		panic(err)
	}

	// 1. getting tiles from the image
	// (img, width, height, overlap)
	tiles, err := GetTilesFromImg(img, 448, 448, 50)
	if err != nil {
		panic(err)
	}

	// 2. consolidating the tiles
	images := make([]api.ImageData, 0)
	for _, v := range tiles {
		buf := new(bytes.Buffer)
		if err := png.Encode(buf, v); err != nil {
			panic(err)
		}
		images = append(images, api.ImageData(buf.Bytes()))
	}

	// 3. adding tiles to the context
	messages := []api.Message{
		// ... omitted for brevity
		{Role: "user",
			Content: "Extract the text from this image exactly as it appears",
			Images:  images,
		},
	}

	// set LLM options
	options := map[string]any{
		"temperature": 0.0,  // force the most likely token
		"top_p":       0.1,  // only consider high-probability tokens
		"num_ctx":     4096, // ensure enough context for the image tokens
	}

	falsePtr := false
	req := &api.ChatRequest{
		// models I tried: "qwen3-vl:235b-cloud, llava:7b, granite3.2-vision:latest, qwen3-vl:2b, bakllava:7b, gemma3:1b, moondream:1.8b"
		Model:    "qwen3-vl:235b-cloud",
		Messages: messages,
		Stream:   &falsePtr,
		Options:  options,
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

	// 5. chatting with Ollama
	scanner := bufio.NewScanner(os.Stdin)
	for {
		// 5.1 get user's prompt
		fmt.Print("\nUser > ")
		if !scanner.Scan() {
			break
		}
		prompt := scanner.Text()
		if strings.TrimSpace(prompt) == "" {
			continue
		}

		// 5.2 append user prompt to chat history
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
