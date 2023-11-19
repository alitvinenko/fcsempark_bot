ENVS=-e DATABASE_PATH='/data/db.json' -e CHAT_ID -e TOKEN
IMAGE_NAME="fcsempark_bot:latest"

build:
	docker build --rm -t ${IMAGE_NAME} .

run-daemon:
	docker run -d -v "/tmp/db.json:/data/db.json" ${ENVS} ${IMAGE_NAME} sh -c './main'

run-createpoll:
	docker run -v "/tmp/db.json:/data/db.json" ${ENVS} ${IMAGE_NAME} sh -c './createpoll'

run-showdb:
	docker run -v "/tmp/db.json:/data/db.json" ${ENVS} ${IMAGE_NAME} sh -c './showdb'