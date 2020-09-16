package p

import (
	"bytes"
	"encoding/json"
	"html"
	"log"
	"net/http"
	"time"
)

type webHookReqBody struct {
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

// sendMessage bla bla
func sendMessage(chatID int64, text string) {

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

	res, err := http.Post("https://api.telegram.org/bot1264325792:AAGhPjE2GOEJoiL4DR_MTrddleuWKuD4KRA/sendMessage", "application/json", bytes.NewBuffer(reqBytes))

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

// GarbageCollect handle the request and sent back the related response
func GarbageCollect(res http.ResponseWriter, req *http.Request) {

	body := &webHookReqBody{}
	if err := json.NewDecoder(req.Body).Decode(body); err != nil {
		log.Println("An error occured (webHookHandler)")
		log.Println(err)
		return
	}

	log.Println(body)

	switch weekday := time.Now().Weekday(); weekday {

	case time.Tuesday, time.Thursday, time.Saturday, time.Wednesday:
		log.Println("Weekday: " + weekday.String())
		if body.Message.Text == "/garbage" || body.Message.Text == "/trash" {
			sendMessage(body.Message.Chat.ID, "# Hello guys! \nToday is **"+weekday.String()+"** and wish you have a great day! \n\nRemember to take out the trash, please. \n ![Tux, the Linux mascot](https://www.iconfinder.com/data/icons/competitive-strategy-and-corporate-training/512/170_waste_disposal_garbage_management_recycle_network_cloud_smart_computing-512.png)")
		}
	}

	log.Println(res, html.EscapeString("Finished..."))
	return
}
