FROM golang:1.21-alpine as builder

WORKDIR /app

COPY ./ /app

RUN go mod download

RUN go test -coverpkg=./... ./... 

RUN CGO_ENABLED=0 go build -o bin/ cmd/api/main.go

EXPOSE 9000

VOLUME [ "./assets" ]
ENTRYPOINT ["./bin/main"]