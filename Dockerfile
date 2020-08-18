FROM golang:1.14.5-buster AS debug
WORKDIR /app
COPY . .
CMD go run gorm-test
