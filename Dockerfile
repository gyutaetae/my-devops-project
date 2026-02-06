FROM golang:1.21-alpine
WORKDIR /app
COPY . .
RUN go mod init my-app && go build -o server .
EXPOSE 8080
CMD ["./server"]
