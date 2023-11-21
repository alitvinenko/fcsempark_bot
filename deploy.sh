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
DOCKER_IMAGE=${DOCKER_REGISTRY}/${IMAGE_NAME}:main

echo `CONTAINER_NAME=${CONTAINER_NAME}`
echo `DOCKER_REGISTRY=${DOCKER_REGISTRY}`
echo `DOCKER_IMAGE=${DOCKER_IMAGE}`

echo "Stop running container"
# Остановка и удаление текущего контейнера
docker stop $CONTAINER_NAME || true
docker rm $CONTAINER_NAME || true

echo "Pull image"
# Загрузка обновленного образа
docker pull $DOCKER_IMAGE

echo "Run container and daemon"
# Запуск нового контейнера с переменными окружения из файла .env
docker run -d --rm --name $CONTAINER_NAME -v "/usr/local/_data:/data" --env-file .env $DOCKER_IMAGE ./main