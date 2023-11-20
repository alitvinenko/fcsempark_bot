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

1. Create directory for database file
```shell
sudo mkdir /usr/local/_data
```

2. Set env variables
```shell
export TOKEN=<your bot token>
export CHAT_ID=<chat id>
```

3. Build image
```shell
make build
```

4. Run bot daemon
```shell
make run-daemon
```

5. Add to crontab schedule when we want to create polls
```shell
# /etc/crontab

0 8 * * sun,thu 
```