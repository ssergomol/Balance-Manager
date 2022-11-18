FROM golang:1.18-alpine

WORKDIR /app

COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY . /app/

RUN go build -o /balance-manager
CMD [ "/balance-manager" ]