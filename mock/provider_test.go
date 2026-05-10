package mock

import (
	"testing"

	"github.com/brian-ding/blaw/llm"
)

func TestChat_ReturnsText(t *testing.T) {
	// arrange
	provider := NewMockProvider([]llm.AssistantTurn{{Text: "Hello, world!"}})

	// act
	response, err := provider.Chat(nil, nil)

	// assert
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if response.Text != "Hello, world!" {
		t.Errorf("expected response text to be 'Hello, world!', got '%s'", response.Text)
	}
}

func TestChat_ReturnsToolCalls(t *testing.T) {
	// arrange
	provider := NewMockProvider([]llm.AssistantTurn{
		{
			Text: "Here are the results:",
			ToolCalls: []llm.ToolCall{
				{Name: "search", Arguments: map[string]string{"query": "golang"}},
				{Name: "calculate", Arguments: map[string]float64{"a": 1, "b": 2}},
			},
		},
	})

	// act
	response, err := provider.Chat(nil, nil)

	// assert
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if len(response.ToolCalls) != 2 {
		t.Errorf("expected 2 tool calls, got %d", len(response.ToolCalls))
	}
	if response.ToolCalls[0].Name != "search" {
		t.Errorf("expected first tool call to be 'search', got '%s'", response.ToolCalls[0].Name)
	}
	if response.ToolCalls[1].Name != "calculate" {
		t.Errorf("expected second tool call to be 'calculate', got '%s'", response.ToolCalls[1].Name)
	}
}

func TestChat_StepsThroughSequence(t *testing.T) {
	// arrange
	provider := NewMockProvider([]llm.AssistantTurn{
		{Text: "First response"},
		{Text: "Second response"},
	})

	// act & assert
	response1, err1 := provider.Chat(nil, nil)
	if err1 != nil {
		t.Fatalf("expected no error, got %v", err1)
	}
	if response1.Text != "First response" {
		t.Errorf("expected first response text to be 'First response', got '%s'", response1.Text)
	}

	response2, err2 := provider.Chat(nil, nil)
	if err2 != nil {
		t.Fatalf("expected no error, got %v", err2)
	}
	if response2.Text != "Second response" {
		t.Errorf("expected second response text to be 'Second response', got '%s'", response2.Text)
	}
}

func TestChat_ReturnsDefaultWhenNoResponsesLeft(t *testing.T) {
	// arrange
	provider := MockProvider{}

	// act
	response, err := provider.Chat(nil, nil)

	// assert
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	expectedText := "MockProvider: no more responses"
	if response.Text != expectedText {
		t.Errorf("expected response text to be '%s', got '%s'", expectedText, response.Text)
	}
}

func TestChat_IgnoresMessagesAndTools(t *testing.T) {
	// arrange
	provider := NewMockProvider([]llm.AssistantTurn{{Text: "Response regardless of input"}})

	// act
	response, err := provider.Chat([]llm.Message{{Content: "Hello"}}, []llm.ToolDefinition{{Name: "tool1"}})

	// assert
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if response.Text != "Response regardless of input" {
		t.Errorf("expected response text to be 'Response regardless of input', got '%s'", response.Text)
	}
}
