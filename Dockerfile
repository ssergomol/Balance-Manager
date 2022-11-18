FROM golang:1.18-alpine

WORKDIR /app

COPY go.mod ./
COPY go.sum ./

RUN go mod download

# COPY ./cmd ./
# COPY ./pkg ./
# COPY ./reports ./
# COPY config.toml ./
COPY . ./

RUN go build -o balance-manager cmd/main.go

EXPOSE 8080

CMD [ "./balance-manager" ]