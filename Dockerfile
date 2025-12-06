FROM golang:alpine AS builder

ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64 \
    GOTOOLCHAIN=auto


WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download


COPY . .


RUN mkdir -p /app/bin && \
    go build -trimpath -ldflags="-s -w" -o /app/bin/app ./cmd/app


FROM gcr.io/distroless/static:nonroot


WORKDIR /app


COPY --from=builder /app/bin/app /app/app

USER nonroot:nonroot

EXPOSE 8080


ENTRYPOINT ["/app/app"]
