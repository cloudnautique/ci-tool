FROM golang:1.4

RUN go get github.com/tools/godep
COPY . /go/src/github.com/cloudnautique/ci-tool
WORKDIR /go/src/github.com/cloudnautique/ci-tool
RUN godep go build -ldflags "-linkmode external -extldflags -static" -o /usr/bin/ci-tool main.go report.go types.go
CMD ["ci-tool"]
