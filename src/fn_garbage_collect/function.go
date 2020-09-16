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
		if body.Message.Text == "/garbage" || body.Message.Text == "/trash" {
			sendTelegramMessage(body.Message.Chat.ID, BuildTelegramMessage(weekday.String()))
		}
	}

	log.Println(res, html.EscapeString("Finished..."))
	return
}
