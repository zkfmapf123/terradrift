FROM golang:1.24-alpine as builder

WORKDIR /app

COPY . . 

RUN go build -o terradrift main.go

FROM scratch as runner

WORKDIR /app

COPY --from=builder /app/terradrift .

ENTRYPOINT ["/app/terradrift"] 



