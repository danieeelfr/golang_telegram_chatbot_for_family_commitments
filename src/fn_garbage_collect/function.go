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
	log.Println("0")
	body := &telegramWebHookReqBody{}
	if err := json.NewDecoder(req.Body).Decode(body); err != nil {
		log.Println("1")
		log.Println(err)
		return
	}
	log.Println("2")
	log.Println(body)
	log.Println("3")
	switch weekday := time.Now().Weekday(); weekday {

	case time.Tuesday, time.Thursday, time.Saturday:
		log.Println("4")
		sendTelegramMessage(body.Message.Chat.ID, BuildTelegramMessage(weekday.String()))
		log.Println("5")
	default:
		if body.Message.Text == "/lixo" || body.Message.Text == "/lixeira" ||
			body.Message.Text == "/rotina" || body.Message.Text == "/tarefas" ||
			body.Message.Text == "/o que fazer" || body.Message.Text == "/lembrar" {
			log.Println("6")
			sendTelegramMessage(body.Message.Chat.ID, BuildTelegramMessage(weekday.String()))
			log.Println("7")
		}
		log.Println("8")
	}
	log.Println("9")

	log.Println(res, html.EscapeString("Finished..."))
	log.Println("10")
	return
}
