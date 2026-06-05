# STAGE 1: Compilation Sandbox
FROM golang:1.22-alpine AS builder
LABEL stage="builder"

WORKDIR /app
COPY go.mod ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-w -s" -o main .

# STAGE 2: Micro-Runtime Environment
FROM alpine:3.19 AS runner
LABEL version="1.0.0" description="ASCII Art Web Runtime Container"

RUN apk --no-cache add ca-certificates && \
    addgroup -S runtimegroup && adduser -S runtimeuser -G runtimegroup

WORKDIR /home/runtimeuser

# Copy built app binary alongside template folders
COPY --from=builder /app/main .
COPY --from=builder /app/templates ./templates
COPY --from=builder /app/static ./static
COPY --from=builder /app/banners ./banners

RUN chown -R runtimeuser:runtimegroup /home/runtimeuser
USER runtimeuser

EXPOSE 8080
CMD ["./main"]