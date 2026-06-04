# Stage 1: Build SPA
FROM node:20-alpine AS spa-builder
WORKDIR /spa
COPY web/package.json web/pnpm-lock.yaml ./
RUN npm ci
COPY web/ .
RUN npm run build

# Stage 2: Build Go binary
FROM golang:1.22-alpine AS go-builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o app ./cmd/app

# Stage 3: Runtime
FROM alpine:3.19
RUN apk --no-cache add ca-certificates tzdata
WORKDIR /app
COPY --from=go-builder /app/app .
COPY --from=spa-builder /spa/public/build ./public/build

EXPOSE 8000
CMD ["./app"]
