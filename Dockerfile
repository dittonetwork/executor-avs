FROM golang:1.21 as builder

ARG APP

RUN --mount=type=secret,id=CI_GITHUB_TOKEN \
    CI_GITHUB_TOKEN=$(cat /run/secrets/CI_GITHUB_TOKEN) && \
    git config --global url.https://ci:${CI_GITHUB_TOKEN}@github.com/.insteadOf https://github.com/

WORKDIR /app
COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o bin/app ./cmd/$APP

FROM alpine:latest

ARG APP

RUN apk add --no-cache tzdata && apk --no-cache add ca-certificates

COPY --from=builder /app/bin/app /bin/app
COPY --from=builder /app/cmd/$APP/config/config.yml /config/

ENTRYPOINT ["/bin/app", "-c", "/config/config.yml"]