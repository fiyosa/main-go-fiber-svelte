FROM node:24-alpine AS spa-builder
WORKDIR /spa
COPY web/ .
RUN pnpm install --frozen-lockfile && pnpm build && rm -rf node_modules

FROM golang:1.22-alpine AS go-builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o app ./cmd/app

FROM alpine:3.19
RUN apk --no-cache add ca-certificates tzdata
WORKDIR /app
COPY --from=go-builder /app/app .
COPY --from=spa-builder /spa/public/build ./public/build
EXPOSE 8000
CMD ["./app"]
