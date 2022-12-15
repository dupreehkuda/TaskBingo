FROM golang:alpine AS builder

WORKDIR /build

COPY . .

RUN go build -o build cmd/main.go

FROM alpine

WORKDIR /build

COPY --from=builder /build /build

CMD ./build
