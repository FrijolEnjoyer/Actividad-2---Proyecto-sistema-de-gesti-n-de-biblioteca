# Build stage
FROM golang:1.22 as builder
WORKDIR /app
COPY . .
RUN go build -o /out/server ./cmd/server

# Runtime stage
FROM gcr.io/distroless/base-debian12
WORKDIR /app
COPY --from=builder /out/server /app/server
ENV PORT=8080
EXPOSE 8080
USER 65532:65532
ENTRYPOINT ["/app/server"]
