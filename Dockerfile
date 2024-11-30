FROM golang:1.23.3-alpine3.20 AS build_base
WORKDIR /app
COPY . .
RUN go get -u -t -d -v ./... && go mod download && go mod tidy && go mod vendor

# Build the Go app
RUN go build -o ./app -v ./cmd/api

FROM alpine:3.20
COPY --from=build_base /app /app
EXPOSE 8081
WORKDIR /app

CMD ["./app"]