FROM golang:alpine as builder
WORKDIR /src
COPY . .
RUN go build -o blog_service .

FROM alpine
WORKDIR /app
COPY --from=builder /src/blog_service /usr/bin/
ENTRYPOINT [ "blog_service" ]
