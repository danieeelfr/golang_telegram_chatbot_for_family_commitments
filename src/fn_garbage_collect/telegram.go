package p

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
)

type telegramWebHookReqBody struct {
	Message struct {
		Text string `json:"text"`
		Chat struct {
			ID int64 `json:"id"`
		} `json:"chat"`
	} `json:"message"`
}

type sendMessageReqBody struct {
	ChatID    int64  `json:"chat_id"`
	Text      string `json:"text"`
	ParseMode string `json:"parse_mode"`
}

func sendTelegramMessage(chatID int64, text string) {

	reqBody := &sendMessageReqBody{
		ChatID:    chatID,
		Text:      text,
		ParseMode: "Markdown",
	}

	reqBytes, err := json.Marshal(reqBody)
	if err != nil {
		log.Println("An error occured (json.Marshal)")
		log.Println(err)
		return
	}

	res, err := http.Post("https://api.telegram.org/bot1264325792:AAEYK9QAWBuoY0U-izjpZQDvzgYaQFwK9mc/sendMessage", "application/json", bytes.NewBuffer(reqBytes))

	if err != nil {
		log.Println("An error occured (calling api with POST)")
		log.Println(err)
		return
	}

	if res.StatusCode != http.StatusOK {
		log.Println(res.Status)
		log.Println(err)
		return
	}

	defer res.Body.Close()

	return
}

// BuildTelegramMessage bla
func BuildTelegramMessage(weekday string) string {
	const EmojiSun string = "\xE2\x98\x80"
	const EmojiRecicle string = "\xE2\x99\xBB"

	greetings := EmojiSun + " Bom dia! " + EmojiSun + "\n\n"
	message := "Hoje é *" + weekday + "*! Espero que você tenha um ótimo dia!\n\n"
	footer := "Por favor, *lembre-se de levar o lixo para fora.* " + EmojiRecicle

	return greetings + message + footer
}
