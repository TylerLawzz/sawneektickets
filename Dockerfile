# Build container
FROM golang:1.22 AS builder

RUN go version

RUN apt-get update && apt-get upgrade -y && apt-get install -y ca-certificates git zlib1g-dev

WORKDIR /go/src/github.com/TicketsBot/worker

# Copy go.mod and go.sum first to leverage caching
COPY go.mod go.sum ./
RUN go mod download && go mod verify

# Now copy the rest of the source files
COPY . .

RUN git submodule update --init --recursive --remote

RUN GOOS=linux GOARCH=amd64 \
    go build \
    -tags=jsoniter \
    -trimpath \
    -o main cmd/worker/main.go

# Prod container
FROM ubuntu:latest

RUN apt-get update && apt-get upgrade -y && apt-get install -y ca-certificates curl

COPY --from=builder /go/src/github.com/TicketsBot/worker/locale /srv/worker/locale
COPY --from=builder /go/src/github.com/TicketsBot/worker/main /srv/worker/main

RUN chmod +x /srv/worker/main

RUN useradd -m container
USER container
WORKDIR /srv/worker

CMD ["/srv/worker/main"]
