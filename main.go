package main

import (
	"bufio"
	"encoding/json"
	"log"
	"os"

	"github.com/benodiwal/lspls/analysis"
	"github.com/benodiwal/lspls/lsp"
	"github.com/benodiwal/lspls/rpc"
)

func main() {
	logger := getLogger("/home/user/personal/lspls/log.txt")
	logger.Println("Hey there!!")

	state := analysis.NewState()

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(rpc.Split)

	for scanner.Scan() {
		msg := scanner.Bytes()
		method, contents,  err := rpc.DecodeMessage(msg)
		if err != nil {
			logger.Printf("Got an error: %s", err)
			continue
		}
		handleMessage(logger, state, method, contents)
	}
}

func handleMessage(logger *log.Logger, state analysis.State, method string, contents []byte) {
	logger.Printf("Received msg with method: %s", method)

	switch method {
		case "initialize":
			var request lsp.InitializeRequest
			if err := json.Unmarshal(contents, &request); err != nil {
				logger.Printf("Hey, we couldn't parse this: %s", err)
			}
			logger.Printf("Connected to: %s %s", request.Params.ClientInfo.Name, request.Params.ClientInfo.Version)

			// Reply
			msg := lsp.NewInitializeResponse(request.ID)
			reply := rpc.EncodeMessage(msg)

			writer := os.Stdout
			writer.Write([]byte(reply))

			logger.Printf("Sent the reply")

		case "textDocument/didOpen":
			var request lsp.DidOpenTextDocumentNotification
			if err := json.Unmarshal(contents, &request); err != nil {
				logger.Printf("textDocument/didOpen: %s", err)
				return
			}
			logger.Printf("Opened: %s", request.Params.TextDocument.URI)
			state.OpenDocument(request.Params.TextDocument.URI, request.Params.TextDocument.Text)
		
		case "textDocument/didChange":
			var request lsp.DidChangeTextDocumentNotification
			if err := json.Unmarshal(contents, &request); err != nil {
				logger.Printf("textDocument/didChange: %s", err)
				return
			}
			logger.Printf("Changed: %s", request.Params.TextDocument.URI)
			for _, change := range request.Params.ContentChanges {
				state.UpdateDocument(request.Params.TextDocument.URI, change.Text)
			}
	}
} 

func getLogger(filename string) *log.Logger {
	logfile, err := os.OpenFile(filename, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0666)
	if err != nil {
		panic("hey, you didn't give me a good file")
	}

	return log.New(logfile, "[lsp]", log.Ldate|log.Ltime|log.Lshortfile)
}