# Development / builder image
FROM golang:1.14-alpine as development
RUN apk add --no-cache make

WORKDIR /app
COPY . .

RUN make build-linux

ENV PORT 8080
EXPOSE 8080
CMD ["/app/cmd/toast/toast"]

# Production image
FROM alpine
WORKDIR /app
COPY --from=development /app/cmd/toast .

ENV PORT 8080
EXPOSE 8080
CMD ["/app/toast"]
