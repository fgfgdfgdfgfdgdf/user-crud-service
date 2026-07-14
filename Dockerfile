FROM golang:1.26.5-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -trimpath -ldflags="-s -w" -o main ./cmd

FROM alpine:3.20

RUN apk --no-cache add ca-certificates
WORKDIR /app

COPY --from=arigaio/atlas:latest /atlas /usr/local/bin/atlas
COPY --from=builder /app/main .
COPY atlas.hcl docker/entrypoint.sh ./
COPY internal/infra/db/schema.sql internal/infra/db/schema.sql
COPY internal/infra/db/migrations internal/infra/db/migrations

RUN chmod +x /app/entrypoint.sh

ENTRYPOINT ["/app/entrypoint.sh"]
CMD ["./main"]