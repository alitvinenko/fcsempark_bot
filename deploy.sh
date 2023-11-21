#!/bin/bash

# Проверка наличия файла .env
if [ ! -f .env ]; then
  echo "Error: .env file not found!"
  exit 1
fi

# Загрузка переменных окружения из файла .env
source .env

# Параметры
CONTAINER_NAME=${CONTAINER_NAME:-fcsempark_bot_daemon}
DOCKER_REGISTRY=${DOCKER_REGISTRY:-your-docker-registry}
DOCKER_IMAGE=${DOCKER_REGISTRY}/${CONTAINER_NAME}:latest

# Остановка и удаление текущего контейнера
docker stop $CONTAINER_NAME || true
docker rm $CONTAINER_NAME || true

# Загрузка обновленного образа
docker pull $DOCKER_IMAGE

# Запуск нового контейнера с переменными окружения из файла .env
docker run -d --name $CONTAINER_NAME -p 8080:8080 --env-file .env $DOCKER_IMAGE ./main