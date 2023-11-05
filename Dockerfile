FROM golang:1.21-alpine AS build

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o rss-confluence-feeder

FROM alpine:latest

WORKDIR /app

COPY --from=build /app/.env .

COPY --from=build /app/rss-confluence-feeder ./rss-confluence-feeder

CMD ["/app/rss-confluence-feeder"]
