build:
	go build -o ./.bin/bot cmd/bot/main.go

run: build
	./.bin/bot

build-image:
	docker build -t cz-api:v0.1 .

start-container:
	docker run --name cz-api -p 80:80 cz-api:v0.1

start-nginx:
	docker compose -f nginx-proxy-compose.yaml up -d