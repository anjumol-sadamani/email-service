FROM golang:latest

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o email-service .

RUN chmod +x email-service

ENTRYPOINT [ "/app/email-service" ]