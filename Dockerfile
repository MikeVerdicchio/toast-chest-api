FROM golang:1.14-alpine

WORKDIR /app
COPY cmd ./cmd

RUN GOFLAGS=-mod=vendor GOOS=linux GOARCH=amd64 CGO_ENABLED=0 \
    cd cmd/toast && go build -o toast main.go

EXPOSE 8080
CMD ["/app/cmd/toast/toast"]
