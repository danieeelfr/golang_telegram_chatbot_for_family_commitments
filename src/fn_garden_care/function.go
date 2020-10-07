package p

import (
	"encoding/json"
	"log"
	"net/http"
)

// GardenCare handle the request and sent back the related response
func GardenCare(res http.ResponseWriter, req *http.Request) {

	body := &telegramWebHookReqBody{}
	if err := json.NewDecoder(req.Body).Decode(body); err != nil {
		log.Println(err)
		return
	}

	log.Println(body)

	if body.Message.Text == "/plantas" || body.Message.Text == "/jardim" ||
		body.Message.Text == "/natureza" || body.Message.Text == "/tarefas" ||
		body.Message.Text == "/o que fazer" || body.Message.Text == "/lembrar" {
		sendTelegramMessage(body.Message.Chat.ID, BuildTelegramMessage())
	}

	return
}
