FROM golang:1.22 as builder

WORKDIR /app
COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .

RUN make build BUILD_MODE=release

FROM alpine:latest

COPY --from=builder /app/bin/operator /bin/operator

ENTRYPOINT ["/bin/operator"]
