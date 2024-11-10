FROM golang:1.23.3-alpine
WORKDIR /app
COPY go.* .
RUN go mod download
COPY main.go .
COPY **/*.go .
RUN go build -o full-stack-crud main.go
RUN chmod +x full-stack-crud
EXPOSE 8080
CMD ["./full-stack-crud"]
