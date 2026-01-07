# PowerPoint Slide Creation Guide

## Slide 1: Title Slide

Title: OCR Magic in Go: Build AI-Powered Text Extraction
Subtitle: ossan, March 2025 Conference
Notes: Start with a joke: "I thought extracting text from a PDF would be simple. Famous last words."

## Slide 2: Speaker Introduction

Title: Who is ossan?
Content:

- Go developer
- Photo of me looking tired
Notes: Joke: "I thought I was a Go expert until I tried AI"

## Slide 3: Today's Mission

Title: What are we doing today?
Content:

- Extract text from PDFs
- Masked as PNGs
- With AI
Notes: Joke: "How hard could it be?"

## Slide 4: The Goal

Title: Our Challenge
Content:

- Query PDFs for information
- From masked images
- Using OCR + VLMs
Notes: Simple diagram: PDF → PNG → AI → Text; Explain the core problem we're solving

## Slide 5: Tech Stack

Title: Technologies Used
Content:

- Go
- Ollama
- Vision Language Models (VLMs)
Notes: Brief explanation of each technology; Why these specific tools

## Slide 6: Why Go?

Title: Go for AI Workflows
Content:

- Fast
- Concurrency
- Lightweight
- Good API integration
Notes: Performance benefits for API-heavy tasks; Memory efficiency (important for constrained hardware)

## Slide 7: Why VLMs?

Title: Vision Language Models
Content:

- Not just OCR
- Context understanding
- Layout awareness
Notes: Joke: "Traditional OCR is like reading with a blindfold"; Explain difference between traditional OCR and VLMs

## Slide 8: Milestone 1 - Local LLM

Title: Starting Simple
Content:

- Local LLM
- Generate endpoint
- Single image
Notes: Joke: "Naive oLocaptimism"; Show the basic approach

## Slide 9: The Code

Title: Basic Implementation
Content:

- client.Generate(ctx, req, respFunc)
Notes: Explain the basic API call; Why this approach was limited

## Slide 10: Results?

Title: Early Results
Content:

- Worked on simple PDFs
- Failed on complex ones
Notes: Joke: "Success is relative"; Show examples of what worked vs. didn't work

## Slide 11: Milestone 2 - Cloud LLM

Title: Scaling Up
Content:

- Cloud LLM
- Same generate endpoint
- Better results
Notes: Joke: "Money solves problems"; Explain why moved to cloud

## Slide 12: Cloud Advantages

Title: Benefits of Cloud Models
Content:

- More powerful models
- Better context
- Larger image support
- Higher cost
Notes: Trade-offs of cloud vs local

## Slide 13: Demo 1

Title: Simple Extraction Demo
Content: Simple extraction
Notes: 2 min recorded demo; Show basic functionality working

## Slide 14: Milestone 3 - Chat Endpoint

Title: Interactive Approach
Content:

- Chat endpoint
- Interactive queries
Notes: Joke: "I wanted a conversation"; Explain why moved to chat endpoint

## Slide 15: New Problems

Title: Unexpected Issues
Content: Images too big
Notes: Joke: "My computer is tiny"; Introduce the image size problem

## Slide 16: Image Size Issue

Title: The Core Problem
Content:

- Large PNGs don't work
- Memory constraints
- Context window limits
Notes: Explain technical limitations; How this affected our approach

## Slide 17: Ollama Logs

Title: Debugging Challenges
Content: Not verbose enough
Notes: Joke: "Logs are like a silent friend"; Explain debugging difficulties

## Slide 18: Tweaking Settings

Title: Configuration Journey
Content:

- Ollama configuration
- Model parameters
Notes: Joke: "Configuration archaeology"; Show what settings were adjusted

## Slide 19: Hallucination Problem

Title: Accuracy Issues
Content:

- AI making things up
- Not literal OCR
Notes: Joke: "Creative AI, not accurate AI"; Explain hallucination problem in OCR

## Slide 20: Model Options

Title: Parameter Tuning
Content:

- Temperature: 0.0
- Top_p: 0.1
- Num_ctx: 4096
Notes: Joke: "The magic numbers"; Explain what each parameter does

## Slide 21: Prompt Engineering

Title: Getting Accurate Results
Content: System prompt: "Be a literal OCR engine"
Notes: Joke: "Teaching AI to be boring"; Explain importance of good prompts

## Slide 22: Cropping Images

Title: Size Reduction Strategy
Content:

- Focus on sections
- Smaller inputs
Notes: Joke: "If at first you don't succeed, crop and try again"; Show before/after examples

## Slide 23: Image Resize

Title: Technical Approach
Content:

- 3rd party Go packages
- Size reduction
- Quality trade-offs
Notes: Mention specific Go packages used; Explain quality vs. performance trade-offs

## Slide 24: Demo 2

Title: Image Processing Demo
Content: Before/after image processing
Notes: 2 min recorded demo; Show image processing improvements

## Slide 25: Reality Check

Title: Hardware Limitations
Content:

- Still couldn't read complex PDFs
- Constrained hardware
Notes: Joke: "My laptop was not impressed"; Emphasize hardware constraints

## Slide 26: The Tiling Idea

Title: Breaking Down the Problem
Content:

- Break large images
- Into smaller tiles
Notes: Joke: "Jigsaw puzzle approach"; Introduce the tiling solution

## Slide 27: Tiling Implementation

Title: Technical Details
Content:

- 448x448 tiles
- 50px overlap
- Go image processing
Notes: Explain the technical approach; Why these specific dimensions

## Slide 28: Tile Processing

Title: Processing Strategy
Content:

- Each tile separately
- Combine results
Notes: Joke: "Like a very expensive 拼图"; Explain the process flow

## Slide 29: Go Tiling Code

Title: Implementation Details
Content: func GetTilesFromImg(r io.Reader, tileWidth, tileHeight, overlap int)
Notes: Show actual code implementation; Explain key parts of the function

## Slide 30: Cloud + Tiling

Title: Best of Both Worlds
Content: Best of both worlds

- Large image support
Notes: Joke: "Finally, a working solution"; Explain how tiling + cloud solved the problem

## Slide 31: Demo 3

Title: Tiling Approach Demo
Content: Tiling approach
Notes: 2 min recorded demo; Show tiling in action

## Slide 32: Hardware Reality

Title: Constrained Environment
Content:

- Constrained machine
- Slow processing
Notes: Joke: "My laptop fan became my alarm clock"; Emphasize the challenges faced

## Slide 33: Performance Issues

Title: Technical Constraints
Content:

- Disk speed
- Memory limits
- CPU constraints
- Temperature problems
Notes: Detail each performance issue; How it affected development

## Slide 34: Workarounds

Title: Creative Solutions
Content:

- External HDD for models
- Performance mode
- Environment tuning
Notes: Joke: "MacGyver coding"; Show creative solutions found

## Slide 35: The Architecture

Title: Final Solution
Content: Go client → Tiled images → Cloud VLMs
Notes: Show architecture diagram; Explain how components work together

## Slide 36: Demo 4

Title: Complete Solution Demo
Content: Complete solution
Notes: 2 min recorded demo; Show full working system

## Slide 37: Lessons Learned

Title: Key Takeaways
Content:

- Local models: constrained
- Image preprocessing: crucial
- Tiling: effective workaround
- Hardware matters
Notes: Summarize key learning points; Emphasize importance of hardware

## Slide 38: Mistakes Made

Title: Learning Through Failure
Content:

- Assumed simple would work
- Ignored image size
- Underestimated hardware needs
Notes: Joke: "Learning through suffering"; Share specific mistakes and lessons

## Slide 39: Technical Challenges

Title: Problem Solving
Content:

- Context windows
- Memory management
- Image formats
- API limitations
Notes: Detail technical challenges faced; How they were addressed

## Slide 40: Success Metrics

Title: What We Achieved
Content:

- Complex PDFs readable
- Accurate extraction
- Scalable approach
- Happy AI (no hallucinations)
Notes: Quantify the success; Show improvement over initial attempts

## Slide 41: Demo 5

Title: Real-World Example Demo
Content: Real-world example
Notes: 2 min recorded demo; Show practical application

## Slide 42: Future Work

Title: Next Steps
Content:

- Real-time processing
- Better error handling
- Model fine-tuning
- Performance optimization
Notes: Outline future improvements; Areas for further development

## Slide 43: Key Takeaway

Title: Main Message
Content:

- AI OCR is complex
- Go handles it well
- Hardware matters
Notes: Joke: "Always test on real hardware"; Summarize the core message

## Slide 44: Thank You

Title: Contact & Resources
Content:

- GitHub: pdf-chatbot
- Sessionize: link
- Questions?
Notes: Joke: "I survived the journey, you can too"; Provide resources for follow-up

## Slide 45: Questions

Title: Q&A Time
Notes: Joke: "This is where I hope you don't ask the hard questions"; Open for questions from audience
