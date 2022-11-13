FROM golang:1.19 as builder
WORKDIR /src/app
COPY . .
RUN GDOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o server -ldflags="-w -s" ./cmd/consumer/main.go

FROM scratch
COPY --from=builder /app/server /server