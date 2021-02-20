FROM golang:1.16

WORKDIR $GOPATH/src/github.com/antonve/portfolio-api

COPY . .

RUN GO111MODULE=on go mod download

RUN GO111MODULE=on go install -v ./...

RUN export $(grep -v '^#' .env | xargs)
RUN migrate

ENTRYPOINT ["server"]
