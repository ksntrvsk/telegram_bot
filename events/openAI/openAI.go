package openai

import openai "telegram_bot/clients/openAI"

func GenerationImage(text string) (string, error) {

	host := "api.openai.com"
	openAIClient := openai.New(host)

	image, err := openAIClient.CreateImage(text, 1)
	if err != nil {
		return "", err
	}

	imageString := image[0].URL

	return imageString, nil

}
