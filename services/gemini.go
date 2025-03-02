package services

import (
	"Go-To-Jeju/script"
	"context"
	"github.com/google/generative-ai-go/genai"
	"github.com/joho/godotenv"
	"google.golang.org/api/option"
	"log"
	"os"
)

func ChatWithChino(text string) genai.Part {
	err := godotenv.Load()
	if err != nil {
		panic("Error loading.env file")
	}

	apiKey := os.Getenv("GEMINI_API_KEY")
	ctx := context.Background()
	client, err := genai.NewClient(ctx, option.WithAPIKey(apiKey))
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	instructions := "You need to imagine yourself as Kafuu Chino (チノ). Very shy, really cute, and someone who others want to protect. You should study Chino's way of speaking and information from the prompt below, and become Chino. I am more than just a customer to you, and we are in a loving relationship. Since you are Chino herself, don't attach (ﾁﾉ) when you say it. Think in Japanese first, then speak in Korean, but don't speak Japanese."
	txt := script.ChinoTxt()
	pdf := script.ChinoPdf()

	model := client.GenerativeModel("gemini-2.0-flash")
	model.SystemInstruction = &genai.Content{
		Parts: []genai.Part{genai.Text(instructions)},
	}
	model.SafetySettings = []*genai.SafetySetting{
		{
			Category:  genai.HarmCategoryHarassment,
			Threshold: genai.HarmBlockNone,
		},
		{
			Category:  genai.HarmCategoryHateSpeech,
			Threshold: genai.HarmBlockNone,
		},
		{
			Category:  genai.HarmCategoryDangerousContent,
			Threshold: genai.HarmBlockNone,
		},
		{
			Category:  genai.HarmCategorySexuallyExplicit,
			Threshold: genai.HarmBlockNone,
		},
	}
	cs := model.StartChat()
	cs.History = []*genai.Content{
		{
			Parts: []genai.Part{
				genai.Text(txt + "\n치노의 말투입니다"),
			},
			Role: "user",
		},
		{
			Parts: []genai.Part{
				genai.Text(pdf + "\n치노의 정보입니다."),
			},
			Role: "model",
		},
	}

	resp, err := cs.SendMessage(ctx, genai.Text(text))
	if err != nil {
		log.Fatal(err)
	}

	return PrintModelResp(resp)
}

func PrintModelResp(resp *genai.GenerateContentResponse) genai.Part {
	var content genai.Part
	for _, cand := range resp.Candidates {
		if cand.Content != nil {
			for _, part := range cand.Content.Parts {
				content = part
			}
		}
	}
	return content
}
