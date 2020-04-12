# Development / builder image
FROM golang:1.14-alpine as development

WORKDIR /app
COPY . .

RUN GOFLAGS=-mod=vendor GOOS=linux GOARCH=amd64 CGO_ENABLED=0 \
    cd cmd/toast && go build -o toast main.go

ENV LISTEN_PORT 8080
EXPOSE 8080
CMD ["/app/cmd/toast/toast"]

# Production image
FROM alpine
WORKDIR /app
COPY --from=development /app/cmd/toast .

ENV LISTEN_PORT 8080
EXPOSE 8080
CMD ["/app/toast"]
