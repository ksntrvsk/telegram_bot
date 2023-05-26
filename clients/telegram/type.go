package telegram

type UpdatesResponce struct {
	Ok     bool     `json:"ok"`
	Result []Update `json:"result"`
}

type Update struct {
	ID      int      `json:"update_id"`
	Message *Message `json:"message"`
}

type Message struct {
	From  From   `json:"from"`
	Chat  Chat   `json:"chat"`
	Text  string `json:"text"`
	Photo string `json:"photo"`
}

type From struct {
	UserName string `json:"username"`
}

type Chat struct {
	ID int `json:"id"`
}
