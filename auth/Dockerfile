FROM golang:alpine as builder
WORKDIR /src
COPY . .
RUN go build -o auth_service .

FROM alpine
WORKDIR /app
COPY --from=builder /src/auth_service /usr/bin/
ENTRYPOINT [ "auth_service" ]
