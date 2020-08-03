# Golang base image
FROM golang:1.14 as go_builder

LABEL stage="tc_builder"

WORKDIR /go/src/github.com/rl404/tax-calculator

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -mod vendor -o tax


# New stage from scratch
FROM alpine:3.10

LABEL maintainer="axel.rl.404@gmail.com"

WORKDIR /docker/bin

COPY --from=go_builder /go/src/github.com/rl404/tax-calculator/tax tax

CMD ["/docker/bin/tax"]