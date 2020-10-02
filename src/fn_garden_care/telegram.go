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

	res, err := http.Post("https://api.telegram.org/bot1264325792:AAGkYKvkGwBbKkwRw_0aHd2X7CEywD_zYls/sendMessage", "application/json", bytes.NewBuffer(reqBytes))

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
func BuildTelegramMessage() string {
	const EmojiSun string = "\xE2\x98\x80"
	const EmojiPlant1 string = "\xF0\x9F\x8C\xB1 \xF0\x9F\x8D\x81"
	const EmojiPlant2 string = "\xF0\x9F\x8C\xB5 \xF0\x9F\x8C\xBE"
	const EmojiPlant3 string = "\n\n\xF0\x9F\x8C\xBB \xF0\x9F\x8C\xBF \xF0\x9F\x8D\x83 \xF0\x9F\x8C\xB4 \xF0\x9F\x8C\xB8"

	greetings := EmojiSun + EmojiPlant1 + " Olá! " + EmojiPlant2 + EmojiSun + "\n\n"
	message := "Vamos aproveitar o dia para *cuidar da natureza?*\n\n"
	footer := "Por favor, *lembre-se de regar e cuidar das plantas do apartamento.* " + EmojiPlant3

	return greetings + message + footer
}
