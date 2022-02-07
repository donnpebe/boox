FROM golang:1.17 as base

FROM base as dev

RUN curl -sSfL https://raw.githubusercontent.com/cosmtrek/air/master/install.sh | sh -s -- -b $(go env GOPATH)/bin

WORKDIR /opt/app/boox
CMD ["air"]

FROM base as built

WORKDIR /opt/app/boox
COPY . .
ENV CGO_ENABLED=0

RUN go get -d -v ./...
RUN go build -o /tmp/boox cmd/main.go

FROM busybox

COPY --from=built /tmp/boox /usr/bin/boox
CMD ["boox"]