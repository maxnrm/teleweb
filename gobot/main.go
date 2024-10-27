package main

import (
	"fmt"
	"log/slog"
	"os"
	"time"

	"gobot/internal/envs"

	tele "gopkg.in/telebot.v3"
)

func main() {
	logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.Level(envs.LOG_LEVEL),
	}))

	slog.SetDefault(logger)

	slog.Info("Starting bot...")

	botSettings := tele.Settings{
		Token:  envs.BOT_TOKEN,
		Poller: &tele.LongPoller{Timeout: 10 * time.Second},
	}

	b, err := tele.NewBot(botSettings)
	if err != nil {
		slog.Error(fmt.Sprintf("bot start error: %s", err.Error()))
		return
	}

	slog.Info("Bot started!")

	b.Handle("/id", Id)

	b.Start()
}

func Id(c tele.Context) error {

	id := fmt.Sprint(c.Chat().ID)

	return c.Send(id)
}
