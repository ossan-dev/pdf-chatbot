single-extract-from-img:
	go build -o app cmd/single-extract-from-img/main.go
	
multiple-extract-from-img:
	go build -o app cmd/multiple-extract-from-img/main.go

chat-ollama-cloud:
	go build -o app cmd/chat-ollama-cloud/main.go