package main

import (
	_ "embed"

	toolset2 "github.com/yeeaiclub/github-a2a/server/toolset"
	"github.com/yeeaiclub/github-a2a/types"
)

//go:embed prompt/system.txt
var systemPrompt string

// AgentConfig represents the configuration for an agent
type AgentConfig struct {
	Tools        map[string]types.Function `json:"tools"`
	SystemPrompt string                    `json:"system_prompt"`
}

// GithubAgent creates a GitHub agent with its tools
func GithubAgent() *AgentConfig {
	toolset := toolset2.NewGitHubToolset()
	tools := toolset.GetTools()

	return &AgentConfig{
		Tools:        tools,
		SystemPrompt: systemPrompt,
	}
}
