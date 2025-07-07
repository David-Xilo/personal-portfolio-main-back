FROM golang:1.24-alpine AS builder

WORKDIR /app

RUN apk --no-cache add ca-certificates git

COPY go.mod go.sum ./
RUN go mod download

COPY . .

ARG BUILD_TIME
ARG GIT_COMMIT
ARG ENV=production

RUN CGO_ENABLED=0 GOOS=linux go build \
    -ldflags="-w -s -X main.BuildTime=${BUILD_TIME} -X main.GitCommit=${GIT_COMMIT} -X main.Environment=${ENV}" \
    -o main .

FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /app

COPY --from=builder /app/main .

ARG ENV=production
ARG GCP_PROJECT_ID

ENV ENV=${ENV}
ENV GCP_PROJECT_ID=${GCP_PROJECT_ID}

ENV PORT=4000
EXPOSE 4000


RUN addgroup -g 1001 -S appgroup && \
    adduser -S appuser -u 1001 -G appgroup

USER appuser

# Health check
#HEALTHCHECK --interval=30s --timeout=3s --start-period=5s --retries=3 \
#    CMD wget --no-verbose --tries=1 --spider http://localhost:4000/health || exit 1

# Start the application
CMD ["./main"]
