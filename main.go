package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {
	http.HandleFunc("/", handleRequest)
	http.ListenAndServe(":8080", nil)
}

func handleRequest(w http.ResponseWriter, r *http.Request) {

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error reading request body", http.StatusBadRequest)
		return
	}

	//  SQL injection attacks using SQLmap
	if strings.Contains(string(body), "--random-agent --technique=T --level=5") {
		sendTelegramMessage(fmt.Sprintf("SQL injection detected from IP: %s with User-Agent: %s", r.RemoteAddr, r.UserAgent()))
		http.Error(w, "SQL injection detected and blocked", http.StatusForbidden)
		return
		//https://github.com/payloadbox/sql-injection-payload-list
	}

	//  XSS attacks using example payloads
	xssPayloads := []string{
		"<script>alert('XSS1');</script>",
		"<img src=\"javascript:alert('XSS2')\">",
		//https://github.com/payloadbox/xss-payload-list
	}
	for _, payload := range xssPayloads {
		if strings.Contains(string(body), payload) {
			sendTelegramMessage(fmt.Sprintf("XSS attack detected from IP: %s with User-Agent: %s", r.RemoteAddr, r.UserAgent()))
			http.Error(w, "XSS attack detected and blocked", http.StatusForbidden)
			return
		}
	}

	// If no attacks detected, pass the request along
	fmt.Fprint(w, "Request successful")
}

func sendTelegramMessage(message string) {
	telegramBotToken := "Telegram_Bot_Token"
	telegramChatID := "Telegram_Chat_ID"

	url := fmt.Sprintf("https://api.telegram.org/bot%s/sendMessage?chat_id=%s&text=%s", telegramBotToken, telegramChatID, message)

	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error sending Telegram message:", err)
	}
	defer resp.Body.Close()
}
