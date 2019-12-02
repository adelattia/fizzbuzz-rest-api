# builder image
FROM golang:latest as builder
WORKDIR /app
COPY . .
# mod vendor => no need to internet connetion to build
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \ 
 go build -mod=vendor -o main .

# run images
FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /app/main .
EXPOSE 8080
CMD ["./main"]