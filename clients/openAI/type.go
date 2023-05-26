package openai

type Response struct {
	Data []URL `json:"data"`
}

type URL struct {
	URL string `json:"url"`
}
