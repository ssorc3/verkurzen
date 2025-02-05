FROM golang:1.23-bookworm

WORKDIR /app

COPY . .

RUN go mod download

RUN go build -o /bin/server ./cmd/server

FROM scratch

COPY --from=0 /bin/server /bin/server
CMD ["/bin/server" "start"]
