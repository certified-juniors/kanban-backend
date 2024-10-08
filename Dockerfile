FROM golang:1.22.4-alpine

RUN mkdir /app
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .

RUN go build -o app ./cmd/main.go

ENTRYPOINT ["./app"]
