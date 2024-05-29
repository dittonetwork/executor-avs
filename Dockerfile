FROM golang:1.21 as builder

ARG APP

RUN --mount=type=secret,id=CI_GITHUB_TOKEN \
    CI_GITHUB_TOKEN=$(cat /run/secrets/CI_GITHUB_TOKEN) && \
    git config --global url.https://ci:${CI_GITHUB_TOKEN}@github.com/.insteadOf https://github.com/ || true

WORKDIR /app
COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .
RUN make go-build-$APP BUILD_MODE=release

FROM alpine:latest

RUN apk add --no-cache tzdata && apk --no-cache add ca-certificates

COPY --from=builder /app/bin/$APP /bin/app

ENTRYPOINT ["/bin/app"]
