package main

import (
	"log"
	"os"
	"time"

	"gopkg.in/telebot.v3"
)

func main() {
	// Отримання токена з середовища
	token := os.Getenv("TELE_TOKEN")
	if token == "" { // Змінено на перевірку на пустоту
		log.Fatal("TELE_TOKEN is not set or empty")
	}
	log.Println("Token found:", token) // Лог для перевірки

	// Створення налаштувань бота
	botSettings := telebot.Settings{
		Token:  token,
		Poller: &telebot.LongPoller{Timeout: 10 * time.Second},
	}

	// Створення об'єкта бота
	bot, err := telebot.NewBot(botSettings)
	if err != nil {
		log.Fatal(err)
	}

	// Обробник текстових повідомлень
	bot.Handle(telebot.OnText, func(c telebot.Context) error {
		return c.Send("Привіт! Як твої справи?")
	})

	// Запуск бота
	bot.Start()
}
