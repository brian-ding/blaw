package llm

type Provider interface {
	Chat(messages []Message, tools []ToolDefinition) (AssistantTurn, error)
}
