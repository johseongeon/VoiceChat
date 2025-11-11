# Build Stage
FROM golang:alpine AS builder

ENV CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

WORKDIR /app

# 먼저 go.mod, go.sum만 복사 → 캐시 활용
COPY go.mod go.sum index.html ./
RUN go mod download

# 이제 나머지 코드 복사
COPY . .

RUN go build -o main cmd/main.go

# Final Stage
FROM scratch

COPY --from=builder /app/main /main
COPY --from=builder /app/index.html /index.html
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

EXPOSE 8080

ENTRYPOINT ["/main"]
