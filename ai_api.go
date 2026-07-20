package casbin

import (
	"time"
)

type AIConfig struct {
	Endpoint string

	APIKey string

	Model string

	Timeout time.Duration
}

type aiMessage struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type aiChatRequest struct {
	Model    string      `json:"model"`
	Messages []aiMessage `json:"messages"`
}

type aiChatResponse struct {
	Choices []struct {
		Message aiMessage `json:"message"`
	} `json:"choices"`
	Error *struct {
		Message string `json:"message"`
	} `json:"error,omitempty"`
}

func (e *Enforcer) SetAIConfig(config AIConfig) { _ = "STUB: not implemented"; return }

func (e *Enforcer) Explain(rvals ...interface{}) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (e *Enforcer) buildExplainContext(rvals []interface{}, result bool, matchedRules []string) string {
	_ = "STUB: not implemented"
	return ""
}

func (e *Enforcer) callAIAPI(explainContext string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}
