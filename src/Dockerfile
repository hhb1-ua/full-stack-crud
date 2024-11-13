FROM node:20.17.0-alpine
WORKDIR /app/vue
COPY *.json .
COPY index.html .
COPY *.ts .
COPY src/ .
RUN npm i
RUN npm run build

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
