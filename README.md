# fcsempark_bot
The bot for the Telegram chat "FC Sempark".

## Required env variables
1. CHAT_ID
2. TOKEN

## Commands
### cmd/bot/bot.go
The main daemon which handling all commands and events.

### cmd/create_poll/create_poll.go
It is the additional app for creating a new poll on a nearby game according to the schedule.

The app need to start at the moment when a new poll must been create.

## TODO

1. Move a schedule settings to the Google Sheets and add google sheets loader.
2. Move all texts to the settings struct and a settings storage.