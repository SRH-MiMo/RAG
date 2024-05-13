package scripts

import (
	"context"
	"github.com/google/generative-ai-go/genai"
	"github.com/joho/godotenv"
	"google.golang.org/api/option"
	"log"
	"os"
)

func RunDream(text string) genai.Part {
	err := godotenv.Load(".env")
	if err != nil {
		panic("Error loading .env file")
	}

	apiKey := os.Getenv("GEMINI_API_KEY")

	ctx := context.Background()
	client, err := genai.NewClient(ctx,
		option.WithAPIKey(apiKey))

	if err != nil {
		log.Fatal(err)
	}

	defer client.Close()

	model := client.GenerativeModel("gemini-pro")
	cs := model.StartChat()
	cs.History = []*genai.Content{
		&genai.Content{
			Parts: []genai.Part{
				genai.Text(text),
			},
			Role: "user",
		},
		&genai.Content{
			Parts: []genai.Part{
				genai.Text("만나서 반갑습니다, 저는 해몽 AI 루나입니다."),
			},
			Role: "model",
		},
	}
	resp, err := cs.SendMessage(ctx, genai.Text(text))
	if err != nil {
		log.Fatal(err)
	}
	str := PrintModelResp(resp)
	return str
}

func PrintModelResp(resp *genai.GenerateContentResponse) genai.Part {
	var content genai.Part
	for _, cand := range resp.Candidates {
		if cand.Content != nil {
			for _, part := range cand.Content.Parts {
				content = part // genai.Part의 Text를 반환합니다.
			}
		}
	}
	return content
}
