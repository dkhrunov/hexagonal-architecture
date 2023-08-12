.PHONY: dc run test lint

dc:
	docker-compose up --remove-orphans --build

run:
	go build -o cmd/port-service/bin/app cmd/port-service/main.go && PORT_SERVICE_HTTP_ADDR=:8080 cmd/port-service/bin/app

test:
	go test -race ./...

lint:
	golangci-lint run