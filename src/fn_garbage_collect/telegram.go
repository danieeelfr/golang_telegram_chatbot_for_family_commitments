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
	log.Println("13")
	reqBody := &sendMessageReqBody{
		ChatID:    -467000473,
		Text:      text,
		ParseMode: "Markdown",
	}

	reqBytes, err := json.Marshal(reqBody)
	if err != nil {
		log.Println("14")
		log.Println("An error occured (json.Marshal)")
		log.Println(err)
		return
	}
	log.Println("15")

	log.Println(reqBytes)
	res, err := http.Post("https://api.telegram.org/bot1157412439:AAFo2ZQVvM11KcX0U_BGKkTQfdv25ul2_Y0/sendMessage", "application/json", bytes.NewBuffer(reqBytes))

	if err != nil {
		log.Println("16")
		log.Println("An error occured (calling api with POST)")
		log.Println(err)
		return
	}

	log.Println("17")

	if res.StatusCode != http.StatusOK {
		log.Println("18")
		log.Println(res.Status)
		log.Println(err)
		return
	}
	log.Println("19")
	defer res.Body.Close()
	log.Println("20")
	return
}

// BuildTelegramMessage bla
func BuildTelegramMessage(weekday string) string {
	log.Println("11")
	const EmojiSun string = "\xE2\x98\x80"
	const EmojiRecicle string = "\xE2\x99\xBB"

	greetings := EmojiSun + " Bom dia! " + EmojiSun + "\n\n"
	message := "Hoje é *" + weekday + "*! Espero que você tenha um ótimo dia!\n\n"
	footer := "Por favor, *lembre-se de levar o lixo para fora.* " + EmojiRecicle
	log.Println("12")
	return greetings + message + footer

}
