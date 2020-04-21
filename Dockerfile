FROM golang:1.14-alpine as builder
RUN apk add --no-cache make

WORKDIR /app
COPY . .

RUN make build-linux

# Production image
FROM alpine
WORKDIR /app
COPY --from=builder /app/cmd/toast .

ENV PORT 8080
EXPOSE 8080
CMD ["/app/toast"]
