FROM golang:1.21.3-bookworm

WORKDIR /app

COPY go.mod ./

RUN go mod download

COPY . ./

RUN go build -o /web ./cmd/web

ENTRYPOINT ["/web"]

