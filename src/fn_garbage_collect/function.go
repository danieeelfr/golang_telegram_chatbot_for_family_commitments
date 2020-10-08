package p

import (
	"encoding/json"
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

	weekday := time.Now().Weekday()

	if body.Message.Text == "/lixo" || body.Message.Text == "/lixeira" ||
		body.Message.Text == "/tarefas" || body.Message.Text == "/lembrar" {
		sendTelegramMessage(body.Message.Chat.ID, BuildTelegramMessage(TranslateWeekdayToPortuguese(weekday.String())))
	}

	return
}

// TranslateWeekdayToPortuguese function that translate bla
func TranslateWeekdayToPortuguese(weekday string) string {

	switch weekday {
	case "Monday":
		return "Segunda-feira"
	case "Tuesday":
		return "Terça-feira"
	case "Wednesday":
		return "Quarta-feira"
	case "Thursday":
		return "Quinta-feira"
	case "Friday":
		return "Sexta-feira"
	case "Saturday":
		return "Sábado"
	case "Sunday":
		return "Domingo"
	default:
		return ""

	}
}
