FROM golang:1.12.4

# set $GOPATH
ENV GOPATH $GOPATH:/go/src

# set app dir
ENV APP_DIR /go/src/github.com/rl404/tax-calculator

RUN apt-get update && \
    apt-get upgrade -y

#revel，revel-cli，gorp，go-sql-driver
RUN go get github.com/revel/revel \
    github.com/revel/cmd/revel \
    github.com/go-gorp/gorp \
    github.com/go-sql-driver/mysql

EXPOSE 9001

WORKDIR ${APP_DIR}