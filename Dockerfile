# Build the Go binary
FROM golang:1.23-alpine AS goapp
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download
RUN go install github.com/a-h/templ/cmd/templ@latest

COPY main.go  .
COPY src/ src/

RUN templ generate
RUN go build -o ./goapp

# Build the final image
FROM alpine:latest as release
COPY --from=goapp /app/goapp /goapp

WORKDIR /app
CMD ["/goapp"]
