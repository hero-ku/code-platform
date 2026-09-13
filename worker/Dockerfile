FROM golang:tip-alpine

WORKDIR /usr/local/app

COPY src ./src
COPY go.mod go.sum ./

CMD ["go", "run", "./src"]
