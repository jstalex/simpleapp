build:
	docker build -t backend:1 --file ./cmd/app/Dockerfile .
	docker build -t migrate:1 --file ./migrate/Dockerfile .
	docker-compose build
run: build
	docker-compose up -d
stop:
	docker-compose down