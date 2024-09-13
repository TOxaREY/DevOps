FROM golang:alpine3.20 as builder
WORKDIR /build
COPY . .
RUN go build -o backend ./cmd/api

FROM alpine:latest
WORKDIR /build
COPY --from=builder /build/backend /build/backend
EXPOSE 8080
CMD ["./backend"]
