FROM golang:1.24.2-alpine

COPY . .

RUN go build -o main ./cmd/service/main.go

CMD ./main 