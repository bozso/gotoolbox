FROM golang:1.15-alpine3.13
WORKDIR /src/gotoolbox

deps:
    COPY go.mod go.sum .
    RUN go mod download
    SAVE ARTIFACT go.mod AS LOCAL go.mod
    SAVE ARTIFACT go.sum AS LOCAL go.sum

build:
    FROM +deps
    COPY . /src/gotoolbox
    ENV CGO_ENABLED 0
    # ENV CGO_LDFLAGS -static
    RUN go build -ldflags "-s -w -extldflags '-static'"
    SAVE ARTIFACT gotoolbox AS LOCAL gotoolbox
