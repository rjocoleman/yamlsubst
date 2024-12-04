FROM golang:1.22-bookworm AS builder

WORKDIR /app

RUN apt-get update && apt-get install -y git

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 go build -ldflags="-w -s" -o /yamlsubst

FROM gcr.io/distroless/static-debian12

COPY --from=builder /yamlsubst /yamlsubst
USER nonroot:nonroot

ENTRYPOINT ["/yamlsubst"]
CMD ["--help"]
