package mock

import "github.com/brian-ding/blaw/llm"

type MockProvider struct {
	response []llm.AssistantTurn
}

func NewMockProvider(response []llm.AssistantTurn) *MockProvider {
	return &MockProvider{response: response}
}

func (p *MockProvider) Chat(messages []llm.Message, tools []llm.ToolDefinition) (llm.AssistantTurn, error) {
	if len(p.response) > 0 {
		resp := p.response[0]
		p.response = p.response[1:]
		return resp, nil
	}

	return llm.AssistantTurn{
		Text: "MockProvider: no more responses"}, nil
}
