# syntax=docker/dockerfile:1

# ==============================================================================
# STAGE 1: Compilation Sandbox
# ==============================================================================
FROM golang:1.26-alpine AS builder
LABEL stage="builder"

WORKDIR /app

# 1. Cache dependencies explicitly
# Docker will reuse this cached layer unless go.mod actually changes
COPY go.mod ./

# 2. Copy the remaining application source files
COPY . .

# 3. Build statically linked binary with optimized size flags
# -w strips DWARF debugging information; -s strips symbol tables
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-w -s" -o main .

# ==============================================================================
# STAGE 2: Micro-Runtime Environment
# ==============================================================================
FROM alpine:3.19 AS runner
LABEL version="1.0.0" description="ASCII Art Web Runtime Container"

# Install security certificates and set up a non-root system user
RUN apk --no-cache add ca-certificates && \
    addgroup -S runtimegroup && adduser -S runtimeuser -G runtimegroup

WORKDIR /home/runtimeuser

# Copy built application assets from the build sandbox stage
COPY --from=builder /app/main .
COPY --from=builder /app/templates ./templates
COPY --from=builder /app/static ./static
COPY --from=builder /app/banners ./banners

# Adjust folder permission contexts to prevent privilege escalation
RUN chown -R runtimeuser:runtimegroup /home/runtimeuser
USER runtimeuser

EXPOSE 8080
CMD ["./main"]