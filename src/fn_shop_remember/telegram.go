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

	log.Println(reqBytes)

	res, err := http.Post("https://api.telegram.org/bot864649786:AAEyjkrWlqnUDnJzxkIz8kxKAFb4tJao7qg/sendMessage", "application/json", bytes.NewBuffer(reqBytes))

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
func BuildTelegramMessage(toRemember string) string {
	const EmojiShop1 string = " \xF0\x9F\x92\xAC	\xF0\x9F\x92\xB0 "

	greetings := "Oi!\n\n"

	message := ""
	footer := ""

	if string(toRemember[0]) == "/" {
		message = "Passando aqui para te lembrar de:" + EmojiShop1 + toRemember + EmojiShop1 + "*\n\n"
		footer = "Caso você queira uma sugestão para lista de compras, responda:\n\n /lista sacolao\n ou \n /lista supermercado."
	} else {
		message = "\xF0\x9F\x92\xAD " + toRemember + " \xF0\x9F\x92\xAD"
	}

	return greetings + message + footer
}
