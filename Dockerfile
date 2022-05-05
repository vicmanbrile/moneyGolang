# syntax=docker/dockerfile:1

FROM golang:1.16-alpine

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . ./

RUN go build -o /moneyGolang

EXPOSE 8080

ENV MONGODB_CONNECTION="mongodb+srv://vicmanbrile:06EpI5YiGzdRaoyD@aplicacioneconomico.4zrhb.mongodb.net/myFirstDatabase?retryWrites=true&w=majority"

CMD ["/moneyGolang"]