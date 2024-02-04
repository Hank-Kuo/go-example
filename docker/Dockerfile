FROM golang:1.21-alpine as builder

ENV HOST localhost
ENV PORT 9000
ENV DEBUG false
ENV POSTGRES_HOST localhost
ENV POSTGRES_USER postgres
ENV POSTGRES_PASSWORD postgres
ENV POSTGRES_DATABASE db
ENV JAEGER_HOST http://localhost:14268/api/traces

WORKDIR /app

COPY ./ /app

RUN go mod download

RUN go test -coverpkg=./... ./... 

RUN CGO_ENABLED=0 go build -o bin/ cmd/api/main.go

EXPOSE 9000

VOLUME [ "./assets" ]
ENTRYPOINT ["./bin/main"]