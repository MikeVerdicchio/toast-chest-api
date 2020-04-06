GOLANGCI := golangci/golangci-lint:v1.24.0

MODULE := github.com/MikeVerdicchio/toast-chest-api
CMD_DIR := toast

# Starts all containers
start:
	@docker-compose up

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

# Cleans up all images
clean: stop
	@docker-compose rm --force

# Collect go modules
vendor:
	@docker-compose run --rm api \
		go mod tidy && go mod vendor

# Formats source code
fmt:
	@docker-compose run --rm api \
		go fmt ./...
