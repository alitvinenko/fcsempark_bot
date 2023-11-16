# fcsempark_bot
The bot for the Telegram chat "FC Sempark".

## Required env variables
1. CHAT_ID
2. TOKEN
3. DATABASE_PATH

## TODO

1. Create Makefile
2. Move a schedule settings to the Google Sheets and add google sheets loader.
2. Move all texts to the settings struct and a settings storage.

## Docker

### Build
```shell
docker build --rm -t <image_name>:<image_tag> .
```

### Run containers

1. Daemon container
```shell
docker run -v "/path/to/db.json:/data/db.json" -e DATABASE_PATH='/data/db.json' -e CHAT_ID="_CHAT_ID_" -e TOKEN="_YOUR_BOT_TOKEN_" <image_name>:<image_tag> sh -c './main'
```

2. Run create new poll task
```shell
docker run -v "/path/to/db.json:/data/db.json" -e DATABASE_PATH='/data/db.json' -e CHAT_ID="_CHAT_ID_" -e TOKEN="_YOUR_BOT_TOKEN_" <image_name>:<image_tag> sh -c './main createpoll'
```

3. Advanced. Show all saved items into Database
```shell
docker run -v "/path/to/db.json:/data/db.json" -e DATABASE_PATH='/data/db.json' -e CHAT_ID="_CHAT_ID_" -e TOKEN="_YOUR_BOT_TOKEN_" <image_name>:<image_tag> sh -c './main showdb'
```