FROM golang:1.23.3-alpine
WORKDIR /app
COPY go.* .
RUN go mod download
COPY *.go .
RUN go build -o full-stack-crud
RUN chmod +x full-stack-crud
COPY app/dist dist
EXPOSE 8080
CMD ["./full-stack-crud"]
