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

// sendTelegramMessage bla bla
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

	res, err := http.Post("https://api.telegram.org/bot1264325792:AAGhPjE2GOEJoiL4DR_MTrddleuWKuD4KRA/sendMessage", "application/json", bytes.NewBuffer(reqBytes))

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

	greetings := EmojiSun + " **Good morning!** " + EmojiSun + "\n\n"
	message := "Today is *" + weekday + "*, and I'd like to wish you a great work day!\n\n"
	footer := "Remember to **take out the " + EmojiRecicle + " **trash** " + EmojiRecicle + "**, please."

	return greetings + message + footer
}
