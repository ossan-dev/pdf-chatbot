# Commands

- `ollama run qwen3-vl:2b`

## Change LLMs local folder

- `sudo nano /etc/systemd/system/ollama.service`
  - add tis line below the `Service` section:
  - `Environment="OLLAMA_MODELS=/mnt/usb-TOSHIBA_EXTERNAL_USB_20210709005169F-0:0-part1/llms"`
- `systemctl daemon-reload`
- `systemctl restart ollama`

## Checks Ollama Logs

- `journalctl -u ollama --no-pager --follow --pager-end`

## Things to Note

- I put the system to run in Performance mode as the CPU power profile
- I benchmarked the disks I have on my machine
  - It turned out that the internal HD is several times more performant of the other external one
  - I'll kept the models save on the external hard disk. I'll be moving them on the internal one when needed

## cURL commands

`BASE64_IMAGE=$(base64 -w 0 ./imgs/sample.png)`
`BASE64_IMAGE=$(base64 -w 0 /home/ossan/Projects/pdf-chatbot/imgs/sample_reduced.png)`

```bash
curl -X POST http://localhost:11434/api/chat -d "{
  \"model\": \"moondream:1.8b\",
  \"messages\": [
    {
      \"role\": \"system\",
      \"content\": \"You are a helpful assistant. Answer by rephrasing the question\"
    },
    {
      \"role\": \"user\",
      \"content\": \"What's the value of the Age field?\",
      \"images\": [\"${BASE64_IMAGE}\"]
    }
  ],
  \"stream\": false
}" | jq '.["message"]["content"]'
```
