ENVS=-e DATABASE_PATH='/data/db.db' -e CHAT_ID -e TOKEN
IMAGE_NAME="fcsempark_bot:latest"
IMAGE="ghcr.io/alitvinenko/fcsempark_bot:main"

build:
	docker build --rm -t ${IMAGE_NAME} .

run-daemon:
	docker run -d --rm --name "fcsempark_bot_daemon" -v "/usr/local/_data:/data" ${ENVS} ${IMAGE_NAME} sh -c './main'

createpoll:
	docker run --rm -v "/usr/local/_data:/data" ${ENVS} ${IMAGE_NAME} sh -c './main createpoll'

showdb:
	docker run --rm -v "/usr/local/_data:/data" ${ENVS} ${IMAGE_NAME} sh -c './main showdb'

createpoll-prod:
	docker run --rm -v "/usr/local/_data:/data" --env-file .env ${IMAGE} ./main createpoll

showdb-prod:
	docker run --rm -v "/usr/local/_data:/data" --env-file .env ${IMAGE} ./main showdb