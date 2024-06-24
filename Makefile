build:
	@echo "Building the application"
	@go build -o build/lspls cmd/main.go

run:
	@echo "Running the application"
	@./build/lspls
