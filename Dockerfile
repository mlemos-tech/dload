FROM golang:1.21-alpine AS build

WORKDIR /app
COPY . /app

RUN go mod download
RUN go build -o ./server ./src/

COPY resource/dev.env ./

EXPOSE 8080

CMD ["./server"]