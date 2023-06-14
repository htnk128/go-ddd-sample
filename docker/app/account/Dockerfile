FROM golang:1.18.10-alpine3.16 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags "-s -w" -trimpath -o bin/account-app cmd/account/server.go

FROM golang:1.18.10-alpine3.16 AS developer

RUN apk update && \
    apk --no-cache add \
    tzdata

# タイムゾーンの設定
RUN cp /usr/share/zoneinfo/Asia/Tokyo /etc/localtime

FROM scratch

# タイムゾーンの設定
COPY --from=developer /usr/share/zoneinfo /usr/share/zoneinfo
COPY --from=developer /usr/local/go/lib/time/zoneinfo.zip /usr/local/go/lib/time/zoneinfo.zip
ENV TZ=Asia/Tokyo

WORKDIR /app

COPY --from=builder /app/bin ./

CMD ["./account-app"]
