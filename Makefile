# Variables
BIN_DIR=bin
BINARY_NAME=file-registry
DOCKER_IMAGE_NAME=$(BINARY_NAME):latest
CONTRACT_DIR=contracts
BUILD_DIR=$(CONTRACT_DIR)/build
BINDINGS_DIR=$(CONTRACT_DIR)/fileregistry

.PHONY: all build run test clean contract env-setup docker-build docker-run docker-up docker-down logs help

all: build test

clean:
	@echo "Cleaning..."
	@rm -rf $(BIN_DIR)
	@rm -rf $(BUILD_DIR)

build:
	@echo "Building $(BINARY_NAME)..."
	@mkdir -p $(BIN_DIR)
	@CGO_ENABLED=1 GOOS=linux go build -ldflags="-w -s" -o $(BIN_DIR)/$(BINARY_NAME) ./cmd/file-registry

run: build
	@echo "Running $(BINARY_NAME)..."
	@./$(BIN_DIR)/$(BINARY_NAME)

test:
	@echo "Running tests..."
	@go test -v ./internal/... ./config/...

test-integration:
	@echo "Running integration tests..."
	@go test -v ./integration/...

env-setup:
	@if [ ! -f .env ]; then \
		echo "Creating .env from template..."; \
		cp .env.template .env; \
	fi

contract:
	@echo "Compiling smart contract..."
	@mkdir -p $(BUILD_DIR)
	@solc --abi --bin $(CONTRACT_DIR)/solidity/hardhat/contracts/FileRegistry.sol -o $(BUILD_DIR) --overwrite
	@echo "Generating Go bindings..."
	@mkdir -p $(BINDINGS_DIR)
	@abigen --bin=$(BUILD_DIR)/FileRegistry.bin \
		--abi=$(BUILD_DIR)/FileRegistry.abi \
		--pkg=fileregistry \
		--type=FileRegistry \
		--out=$(BINDINGS_DIR)/bindings.go

docker-build: build
	@echo "Building Docker image..."
	@docker build -t $(DOCKER_IMAGE_NAME) .

docker-run: docker-build
	@echo "Running Docker container..."
	@docker run -p 8090:8090 $(DOCKER_IMAGE_NAME)

docker-up: env-setup
	@echo "Starting Hardhat service..."
	@docker-compose up --build -d hardhat
	@echo "Waiting for Hardhat node to become ready..."
	@sleep 5
	@echo "Starting remaining services..."
	@docker-compose up --build --remove-orphans

docker-down:
	@echo "Stopping and removing all services..."
	@docker-compose down -v

logs:
	@docker-compose logs -f

help:
	@echo "Usage: make [TARGET]"
	@echo ""
	@echo "Targets:"
	@echo "  all           Build and run tests"
	@echo "  build         Build the application binary"
	@echo "  run           Build and run the application locally"
	@echo "  test          Run tests"
	@echo "  clean         Remove binary and build artifacts"
	@echo "  env-setup     Create .env file from template if it doesn't exist"
	@echo "  contract      Compile contracts and generate Go bindings"
	@echo "  docker-build  Build Docker image"
	@echo "  docker-run    Run Docker container from built image"
	@echo "  docker-up     Start Hardhat, deploy contract, then start other services"
	@echo "  docker-down   Stop and remove all services"
	@echo "  logs          Show docker-compose logs"
	@echo "  help          Show this help message"
