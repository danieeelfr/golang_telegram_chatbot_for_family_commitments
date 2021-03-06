package p

import (
	"encoding/json"
	"log"
	"net/http"
)

// ShopRemember handle the request and sent back the related response
func ShopRemember(res http.ResponseWriter, req *http.Request) {

	body := &telegramWebHookReqBody{}
	if err := json.NewDecoder(req.Body).Decode(body); err != nil {
		log.Println(err)
		return
	}

	if body.Message.Text == "/compras" ||
		body.Message.Text == "/supermercado" || body.Message.Text == "/comprar" ||
		body.Message.Text == "/sacolão" || body.Message.Text == "/tarefas" ||
		body.Message.Text == "/lembrar" {
		sendTelegramMessage(body.Message.Chat.ID, BuildTelegramMessage(body.Message.Text))
	} else if body.Message.Text == "/lista sacolão" {
		list := "- item 1\n"
		list += "- item 2\n"
		list += "- item 3\n"
		list += "- item 4\n"
		sendTelegramMessage(body.Message.Chat.ID, BuildTelegramMessage(list))
	} else if body.Message.Text == "/lista supermercado" {
		list := "- item 1\n"
		list += "- item 2\n"
		list += "- item 3\n"
		list += "- item 4\n"
		list += "- item 1\n"
		list += "- item 2\n"
		list += "- item 3\n"
		list += "- item 4\n"
		sendTelegramMessage(body.Message.Chat.ID, BuildTelegramMessage(list))
	}

	return
}
