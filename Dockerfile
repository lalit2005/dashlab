FROM golang:1.21-alpine
WORKDIR /app
COPY . .
RUN go mod download
RUN go build -o dashlab .
EXPOSE 4000
CMD ["sh", "./startup.sh"]
