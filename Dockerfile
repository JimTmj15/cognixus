FROM golang:1.21.2

WORKDIR /app

COPY go.mod ./
RUN go mod download && go mod verify

COPY . .

RUN go build -o main cmd/main.go

RUN chmod +x /app/main

CMD ["/app/main"]