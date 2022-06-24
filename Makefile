OUTPUT = main 
SERVICE_NAME = products

.PHONY: test
test:
	go test ./...

build-local:
	go build -o $(OUTPUT) ./cmd/$(SERVICE_NAME)/main.go

run: build-local
	@echo ">> Running application ..."
	APP_PORT=9090 \
	./$(OUTPUT)