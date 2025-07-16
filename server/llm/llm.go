package llm

import (
	"context"

	"github.com/cohesion-org/deepseek-go"
	"github.com/yeeaiclub/github-a2a/types"
	"github.com/yumosx/got/pkg/stream"
)

type DeepSeekHandler struct {
	client *deepseek.Client
	model  string
}

func NewDeepSeek(client *deepseek.Client) *DeepSeekHandler {
	return &DeepSeekHandler{client: client}
}

func (d *DeepSeekHandler) Handle(ctx context.Context, message []types.LLMRequest) error {
	completion, err := d.client.CreateChatCompletion(ctx, &deepseek.ChatCompletionRequest{
		Model:    d.model,
		Messages: d.toMessage(message),
	})
	if err != nil {
		return err
	}
	return nil
}

func (d *DeepSeekHandler) toMessage(message []types.LLMRequest) []deepseek.ChatCompletionMessage {
	return stream.Map(message, func(idx int, src types.LLMRequest) deepseek.ChatCompletionMessage {
		return deepseek.ChatCompletionMessage{
			Role:    src.Role,
			Content: src.Content,
		}
	})
}
