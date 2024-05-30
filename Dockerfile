FROM golang:1.21 as builder

ARG APP

WORKDIR /app
COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .

RUN make go-build-$APP BUILD_MODE=release

FROM alpine:latest

ARG APP

RUN apk add --no-cache tzdata && apk --no-cache add ca-certificates

COPY --from=builder /app/bin/$APP /bin/app

ENTRYPOINT ["/bin/app"]
