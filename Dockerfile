FROM golang:1.12.4

# add git repo to container
ADD . /go/src/github.com/rl404/tax-calculator

# import revel，revel-cli，gorp，go-sql-driver
RUN go get github.com/revel/revel \
    github.com/revel/cmd/revel \
    github.com/go-gorp/gorp \
    github.com/go-sql-driver/mysql

# run revel app
ENTRYPOINT revel run github.com/rl404/tax-calculator

EXPOSE 9001
