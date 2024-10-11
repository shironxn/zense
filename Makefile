BIN := main
BIN_DIR := ./bin
TMP_DIR := ./tmp

# ==================================================================================== #
# HELPERS
# ==================================================================================== #

.PHONY: help
help: ## Display usage information
	@echo "Usage:"
	@awk 'BEGIN {FS = ":.*?## "}; /^[a-zA-Z_-]+:.*?## / {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST)

# ==================================================================================== #
# QUALITY CONTROL
# ==================================================================================== #

.PHONY: tidy
tidy: ## Run go mod tidy
	@echo "Running go mod tidy..."
	@go mod tidy

.PHONY: clean
clean: ## Clean the project
	@echo "Cleaning the project..."
	@rm -rf $(BIN_DIR)
	@rm -rf $(TMP_DIR)

# ==================================================================================== #
# APPLICATION
# ==================================================================================== #

.PHONY: test
test: ## Run tests
	@echo "Running tests..."
	@go test -v ./...

.PHONY: dev
dev: ## Run development environment
	@echo "Running development environment..."
	@$(GOPATH)/bin/air --build.cmd "go build -o $(BIN_DIR)/$(BIN) cmd/main.go" --build.bin "$(BIN_DIR)/$(BIN)"

.PHONY: build
build: ## Build the project
	@echo "Building the project..."
	@go build -o $(BIN_DIR)/$(BIN) ./cmd

.PHONY: run
run: tidy build ## Run the project
	@echo "Running the project..."
	@$(BIN_DIR)/$(BIN)

.PHONY: docker-up
docker-up: ## Start Docker Compose services
	@echo "Starting Docker Compose services..."
	@docker-compose up -d

.PHONY: docker-down
docker-down: ## Stop Docker Compose services
	@echo "Stopping Docker Compose services..."
	@docker-compose down
