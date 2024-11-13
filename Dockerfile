FROM golang:1.22.0-alpine as build

WORKDIR /app

COPY go.* ./

RUN go mod download

COPY . ./

RUN go build -o /go-app

FROM alpine:latest

WORKDIR /app

RUN adduser -s /bin/sh -u 1100 --disabled-password  appuser

RUN chown -R appuser:appuser /app

USER appuser

COPY --from=build /go-app /app/go-app

ENTRYPOINT [ "./go-app" ]