.DEFAULT_GOAL=run-app

include gobot/.env
export $(shell sed 's/=.*//' gobot/.env)

.PHONY: run-bot
run-bot:
	cd gobot && go run main.go

.PHONY: run-app
run-app:
	cd phaser-app && bun run dev
