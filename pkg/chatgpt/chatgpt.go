package chatgpt

import (
	"context"
	"fmt"
	"os"

	openai "github.com/sashabaranov/go-openai"
)

type ChatGPTClient struct {
	Client *openai.Client
}

func NewChatGPTClient() *ChatGPTClient {
	API_KEY := os.Getenv("OPENAI_API_KEY")
	c := openai.NewClient(API_KEY)
	return &ChatGPTClient{
		Client: c,
	}
}

func (c ChatGPTClient) Chat(sysMsg, usrMsg string) string {
	ctx := context.Background()
	req := openai.ChatCompletionRequest{
		Model: openai.GPT3Dot5Turbo,
		Messages: []openai.ChatCompletionMessage{
			{
				Role:    openai.ChatMessageRoleSystem,
				Content: sysMsg,
			},
			{
				Role:    openai.ChatMessageRoleUser,
				Content: usrMsg,
			},
		},
	}

	resp, err := c.Client.CreateChatCompletion(ctx, req)
	if err != nil {
		fmt.Printf("ChatCompletion error: %v\n", err)
	}

	return resp.Choices[0].Message.Content
}
