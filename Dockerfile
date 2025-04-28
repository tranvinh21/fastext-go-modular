FROM golang

WORKDIR /app

COPY . .

RUN go build -o main cmd/main.go

EXPOSE 3000

CMD ["./main"]
