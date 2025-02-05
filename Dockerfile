FROM golang:1.23-bookworm AS build

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o server ./cmd/server

FROM alpine:edge

WORKDIR /app

COPY --from=build /app/server .
COPY ./config ./config

RUN apk --no-cache add ca-certificates tzdata

EXPOSE 5000
ENTRYPOINT ["/app/server", "start"]
