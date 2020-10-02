package p

import (
	"encoding/json"
	"html"
	"log"
	"net/http"
	"time"
)

// GarbageCollect handle the request and sent back the related response
func GarbageCollect(res http.ResponseWriter, req *http.Request) {

	body := &telegramWebHookReqBody{}
	if err := json.NewDecoder(req.Body).Decode(body); err != nil {
		log.Println(err)
		return
	}

	log.Println(body)

	switch weekday := time.Now().Weekday(); weekday {

	case time.Tuesday, time.Thursday, time.Saturday:
		sendTelegramMessage(body.Message.Chat.ID, BuildTelegramMessage(weekday.String()))
	default:
		if body.Message.Text == "/lixo" || body.Message.Text == "/lixeira" ||
			body.Message.Text == "/rotina" || body.Message.Text == "/tarefas" ||
			body.Message.Text == "/o que fazer" || body.Message.Text == "/lembrar" {
			sendTelegramMessage(body.Message.Chat.ID, BuildTelegramMessage(weekday.String()))
		}
	}

	log.Println(res, html.EscapeString("Finished..."))
	return
}
