FROM golang:1.18.10-alpine3.16 AS builder

ARG app

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
    go build -tags timetzdata -ldflags "-s -w" -trimpath -o bin/app cmd/${app}/server.go

FROM scratch

ENV TZ=Asia/Tokyo

WORKDIR /app

COPY --from=builder /app/bin ./

CMD ["./app"]
