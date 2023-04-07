tests:
	go test -v ./...

db:
	docker-compose up --build db

mic:
	docker-compose up --build microservice

down:
	docker-compose down

clean:
	docker system prune -f && docker volume prune -f

build:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./zip/bin
