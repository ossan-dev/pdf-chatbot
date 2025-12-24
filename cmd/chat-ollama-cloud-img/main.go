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

type Tile struct {
	Rect image.Rectangle
	Img  image.Image
}

type SubImager interface {
	SubImage(image.Rectangle) image.Image
}

func GetTilesFromImg(r io.Reader, tileWidth, tileHeight, overlap int) ([]Tile, error) {
	tiles := make([]Tile, 0, 1000)
	img, err := png.Decode(r)
	if err != nil {
		return nil, err
	}

	rect := img.Bounds()
	imgWidth := rect.Dx()
	imgHeight := rect.Dy()

	for y := 0; y <= imgHeight; y += tileHeight - overlap {
		for x := 0; x <= imgWidth; x += tileWidth - overlap {
			tileX := x
			tileY := y
			tileHeight := tileHeight
			tileWidth := tileWidth
			if tileX+tileWidth > imgWidth {
				tileWidth = imgWidth - x
			}
			if tileY+tileHeight > imgHeight {
				tileHeight = imgHeight - y
			}
			tileBounds := image.Rect(tileX, tileY, tileX+tileWidth, tileY+tileHeight)
			tileImg := img.(SubImager).SubImage(tileBounds)
			tile := Tile{
				Rect: tileBounds,
				Img:  tileImg,
			}
			tiles = append(tiles, tile)
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

	tiles, err := GetTilesFromImg(imgFile, 448, 448, 50)
	if err != nil {
		panic(err)
	}

	// 3. setup Ollama request

	images := make([]api.ImageData, 0)
	for _, v := range tiles {
		buf := new(bytes.Buffer)
		if err := png.Encode(buf, v.Img); err != nil {
			panic(err)
		}
		images = append(images, api.ImageData(buf.Bytes()))
	}

	// img := resize.Resize(448, 448, image, resize.Lanczos2)

	messages := []api.Message{
		{
			Role: "system",
			Content: "You are a literal OCR engine. Your only job is to extract text from images. " +
				"Rule 1: Never guess or invent data. " +
				"Rule 2: If a field is present but unreadable, write [UNREADABLE]. " +
				"Rule 3: If a field is missing, write [NOT FOUND]. " +
				"Rule 4: Output only the extracted text without any conversational preamble.",
		},
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
