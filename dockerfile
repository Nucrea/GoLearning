FROM golang:1.17-alpine

WORKDIR /app

COPY cmd .
COPY internal .
COPY go.mod .
COPY go.sum .

RUN go mod download

CMD ["go", "run", "cmd/main.go"]