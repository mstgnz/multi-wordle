FROM golang:1.21 AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN GOOS=linux CGO_ENABLED=0 go build -o multiWordle ./cmd

FROM scratch as deploy
COPY --from=builder /app/multiWordle /app/multiWordle
COPY --from=builder /app/public /public
ENTRYPOINT ["/app/multiWordle"]