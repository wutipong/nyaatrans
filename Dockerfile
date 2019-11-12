FROM golang:1.13-alpine

WORKDIR /app

COPY . .

# RUN go test -v .
RUN go build -o main .

CMD ["./main"]
