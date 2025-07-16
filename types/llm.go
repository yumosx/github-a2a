package types

import "github.com/cohesion-org/deepseek-go"

type LLMRequest struct {
	Role    string
	Content string
}

type LLMResponse struct {
	Content string
}

type Function interface {
	FunctionDefinition() deepseek.Function
	Call(args map[string]interface{}) interface{}
}
