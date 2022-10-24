FROM golang:1.16-alpine

WORKDIR /app

COPY cmd .
COPY internal .
COPY go.mod .
COPY go.sum .

CMD ["go", " run", "cmd/main.go"]