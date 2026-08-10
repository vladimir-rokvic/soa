FROM golang:alpine as build

WORKDIR /src

COPY . .

RUN go build -o purchase-service .

FROM alpine

WORKDIR /app

COPY --from=build ./src/purchase-service /usr/bin

ENTRYPOINT [ "purchase-service" ]
