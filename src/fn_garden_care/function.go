package p

import (
	"encoding/json"
	"html"
	"log"
	"net/http"
	"time"
)

// GardenCare handle the request and sent back the related response
func GardenCare(res http.ResponseWriter, req *http.Request) {

	body := &telegramWebHookReqBody{}
	if err := json.NewDecoder(req.Body).Decode(body); err != nil {
		log.Println(err)
		return
	}

	log.Println(body)

	switch weekday := time.Now().Weekday(); weekday {

	case time.Wednesday, time.Monday:
		sendTelegramMessage(body.Message.Chat.ID, BuildTelegramMessage())
	default:
		if body.Message.Text == "/plantas" || body.Message.Text == "/jardim" ||
			body.Message.Text == "/natureza" || body.Message.Text == "/tarefas" ||
			body.Message.Text == "/o que fazer" || body.Message.Text == "/lembrar" {
			sendTelegramMessage(body.Message.Chat.ID, BuildTelegramMessage())
		}
	}

	log.Println(res, html.EscapeString("Finished..."))
	return
}
