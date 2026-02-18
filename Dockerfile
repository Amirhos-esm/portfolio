# ----------------------
# Stage 1: Build
# ----------------------
FROM golang:1.25-alpine AS builder

# Build tools for CGO + SQLite
# RUN apk add --no-cache gcc musl-dev sqlite-dev

WORKDIR /app
RUN apk add --no-cache gcc musl-dev sqlite-dev
# Copy go.mod & go.sum first
COPY go.mod go.sum ./
RUN go mod download

# Copy source
COPY . .

# Build with CGO
RUN CGO_ENABLED=1 go build -o app .

# ----------------------
# Stage 2: Runtime
# ----------------------
FROM alpine:latest

# Required runtime libraries for CGO binaries
# RUN apk add --no-cache libgcc sqlite-libs

WORKDIR /app

# Copy the compiled binary
COPY --from=builder /app/app .

# Optional: copy initial database
# COPY --from=builder /app/*.db .

CMD ["./app"]
