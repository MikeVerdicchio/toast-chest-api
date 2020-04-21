GOLANGCI := golangci/golangci-lint:v1.24.0

MODULE := github.com/MikeVerdicchio/toast-chest-api
CMD_DIR := toast

.PHONY: vendor

# Starts all containers
start:
	@docker-compose up

start-prod:
	@docker-compose -f docker-compose-prod.yml up

# Stops all containers
stop:
	@docker-compose down

# Starts a shell in web container
cli:
	@docker-compose run --rm api ash

# Builds and pulls all images
build:
	@docker-compose pull
	@docker-compose build

build-prod:
	@docker-compose -f docker-compose-prod.yml pull
	@docker-compose -f docker-compose-prod.yml build

# Builds go application for linux
build-linux:
	GOFLAGS=-mod=vendor GOOS=linux GOARCH=amd64 CGO_ENABLED=0 cd cmd/toast && go build -o toast main.go

# Cleans up all images
clean: stop
	@docker-compose rm --force

# Collect go modules
vendor:
	go mod tidy && go mod vendor

# Formats source code
fmt:
	go fmt ./...
