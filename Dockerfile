FROM golang:1-alpine as build

WORKDIR /app
COPY cmd cmd
RUN go build cmd/hello/hello.go

FROM alpine:latest

WORKDIR /app
COPY --from=build /app/hello /app/hello

EXPOSE 80
ENTRYPOINT ["./hello"]