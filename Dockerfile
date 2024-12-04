FROM golang:1.22-alpine AS builder

WORKDIR /app

RUN apk add --no-cache git

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 go build -ldflags="-w -s" -o /yamlsubst

FROM gcr.io/distroless/static:nonroot

COPY --from=builder /yamlsubst /yamlsubst
USER nonroot:nonroot

ENTRYPOINT ["/yamlsubst"]
CMD ["--help"]
