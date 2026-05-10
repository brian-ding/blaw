package llm

type Role int

const (
	System Role = iota
	User
	Assistant
	ToolResult
	Attachment
	Progress
)

type Message struct {
	Id      string `json:"id"`
	Content string `json:"content"`
	Role    Role   `json:"role"`
}

type ToolDefinition struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Parameters  any    `json:"parameters"`
}

type ToolCall struct {
	Id        string `json:"id"`
	Name      string `json:"name"`
	Arguments any    `json:"arguments"`
}

type StopReason int

const (
	Stop StopReason = iota
	ToolUse
)

type TokenUsage struct {
	InputTokens  uint64 `json:"input_tokens"`
	OutputTokens uint64 `json:"output_tokens"`
}

type AssistantTurn struct {
	Text       string     `json:"text"`
	ToolCalls  []ToolCall `json:"tool_calls"`
	StopReason StopReason `json:"stop_reason"`
	Usage      TokenUsage `json:"usage"`
}
