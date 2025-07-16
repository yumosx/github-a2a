package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/yeeaiclub/a2a-go/sdk/client"
	"github.com/yeeaiclub/a2a-go/sdk/types"
)

func main() {
	httpClient := http.Client{}
	newClient := client.NewClient(&httpClient, "http://localhost:8080/api")

	eventChan := make(chan any)
	errChan := make(chan error, 1)

	go func() {
		errChan <- newClient.SendMessageStream(types.MessageSendParam{
			Message: &types.Message{
				ContextID: "context-1",
				Role:      types.User,
				Parts: []types.Part{
					&types.TextPart{Text: "Show recent commits for repository 'facebook/react", Kind: "text"},
				},
			},
		}, eventChan)
		close(eventChan)
	}()

	for event := range eventChan {
		rawMsg, ok := event.(json.RawMessage)
		if !ok {
			log.Println("Unexpected event type")
			continue
		}
		log.Println("Raw event:", string(rawMsg))
	}

	if err := <-errChan; err != nil {
		log.Printf("Stream error: %v", err)
	}
}
