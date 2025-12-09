module pdfchatbot

go 1.25.3

require (
	github.com/nfnt/resize v0.0.0-20180221191011-83c6a9932646
	github.com/ollama/ollama v0.13.3-rc0
)

replace github.com/ollama/ollama => ../ollama

require (
	github.com/google/uuid v1.6.0 // indirect
	golang.org/x/crypto v0.46.0 // indirect
	golang.org/x/sys v0.39.0 // indirect
)
