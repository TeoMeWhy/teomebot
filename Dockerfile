FROM golang:1.22

WORKDIR /app/

COPY . .

RUN go build main.go

CMD ["./main"]