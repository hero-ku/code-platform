FROM golang:tip-alpine

COPY src /src
COPY go.mod /
COPY go.sum /

CMD ["go", "run", "./src"]
