# OCR Magic in Go: Build AI-Powered Text Extraction - 40-Minute Talk Outline

## Slide 1: Title

**OCR Magic in Go: Build AI-Powered Text Extraction**
*ossan*
*March 2025 Conference*

## Slide 2: Me

*Photo of me looking tired*
Go developer
*Joke: "I thought I was a Go expert until I tried AI"*

## Slide 3: Today's Mission

Extract text from PDFs
Masked as PNGs
With AI
*Joke: "How hard could it be?"*

## Slide 4: The Goal

Query PDFs for information
From masked images
Using OCR + VLMs
*Simple diagram: PDF → PNG → AI → Text*

## Slide 5: Tech Stack

- Go
- Ollama
- Vision Language Models (VLMs)

## Slide 6: Why Go?

- Fast
- Concurrency
- Lightweight
- Good API integration

## Slide 7: Why VLMs?

Not just OCR
Context understanding
Layout awareness
*Joke: "Traditional OCR is like reading with a blindfold"*

## Slide 8: Milestone 1

Local LLM
Generate endpoint
Single image
*Joke: "Naive optimism"*

## Slide 9: The Code

```go
client.Generate(ctx, req, respFunc)
```

Simple but limited

## Slide 10: Results?

Worked on simple PDFs
Failed on complex ones
*Joke: "Success is relative"*

## Slide 11: Milestone 2

Cloud LLM
Same generate endpoint
Better results
*Joke: "Money solves problems"*

## Slide 12: Cloud Advantages

- More powerful models
- Better context
- Larger image support
- Higher cost

## Slide 13: Demo 1

Simple extraction
*2 min recorded demo*

## Slide 14: Milestone 3

Chat endpoint
Interactive queries
*Joke: "I wanted a conversation"*

## Slide 15: New Problems

Images too big
*Joke: "My computer is tiny"*

## Slide 16: Image Size Issue

Large PNGs don't work
Memory constraints
Context window limits

## Slide 17: Ollama Logs

Not verbose enough
*Joke: "Logs are like a silent friend"*

## Slide 18: Tweaking Settings

Ollama configuration
Model parameters
*Joke: "Configuration archaeology"*

## Slide 19: Hallucination Problem

AI making things up
Not literal OCR
*Joke: "Creative AI, not accurate AI"*

## Slide 20: Model Options

- Temperature: 0.0
- Top_p: 0.1
- Num_ctx: 4096
*Joke: "The magic numbers"*

## Slide 21: Prompt Engineering

System prompt:
"Be a literal OCR engine"
*Joke: "Teaching AI to be boring"*

## Slide 22: Cropping Images

Focus on sections
Smaller inputs
*Joke: "If at first you don't succeed, crop and try again"*

## Slide 23: Image Resize

3rd party Go packages
Size reduction
Quality trade-offs

## Slide 24: Demo 2

Before/after image processing
*2 min recorded demo*

## Slide 25: The Reality Check

Still couldn't read complex PDFs
Constrained hardware
*Joke: "My laptop was not impressed"*

## Slide 26: The Tiling Idea

Break large images
Into smaller tiles
*Joke: "Jigsaw puzzle approach"*

## Slide 27: Tiling Implementation

- 448x448 tiles
- 50px overlap
- Go image processing

## Slide 28: Tile Processing

Each tile separately
Combine results
*Joke: "Like a very expensive拼图"*

## Slide 29: Go Tiling Code

```go
func GetTilesFromImg(r io.Reader, tileWidth, tileHeight, overlap int)
```

Rectangle calculations

## Slide 30: Cloud + Tiling

Best of both worlds
Large image support
*Joke: "Finally, a working solution"*

## Slide 31: Demo 3

Tiling approach
*2 min recorded demo*

## Slide 32: Hardware Reality

Constrained machine
Slow processing
*Joke: "My laptop fan became my alarm clock"*

## Slide 33: Performance Issues

- Disk speed
- Memory limits
- CPU constraints
- Temperature problems

## Slide 34: Workarounds

- External HDD for models
- Performance mode
- Environment tuning
*Joke: "MacGyver coding"*

## Slide 35: The Architecture

Go client → Tiled images → Cloud VLMs
*Simple diagram*

## Slide 36: Demo 4

Complete solution
*2 min recorded demo*

## Slide 37: Lessons Learned

- Local models: constrained
- Image preprocessing: crucial
- Tiling: effective workaround
- Hardware matters

## Slide 38: Mistakes Made

- Assumed simple would work
- Ignored image size
- Underestimated hardware needs
*Joke: "Learning through suffering"*

## Slide 39: Technical Challenges

- Context windows
- Memory management
- Image formats
- API limitations

## Slide 40: Success Metrics

- Complex PDFs readable
- Accurate extraction
- Scalable approach
- Happy AI (no hallucinations)

## Slide 41: Demo 5

Real-world example
*2 min recorded demo*

## Slide 42: Future Work

- Real-time processing
- Better error handling
- Model fine-tuning
- Performance optimization

## Slide 43: Key Takeaway

AI OCR is complex
Go handles it well
Hardware matters
*Joke: "Always test on real hardware"*

## Slide 44: Thank You

- GitHub: pdf-chatbot
- Sessionize: link
- Questions?
*Joke: "I survived the journey, you can too"*

## Slide 45: Questions

> *Joke: "This is where I hope you don't ask the hard questions"*

---

**Total Slides:** 45
**Total Time:** 40 minutes
**Demo Time:** 10 minutes total (2 min each for 5 demos)
**Speaking Time:** 30 minutes
**Average Time per Slide:** ~53 seconds (allowing for quick transitions)
