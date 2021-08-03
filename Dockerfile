FROM golang:1.16 as build
ARG GO111MODULE=on
ARG CGO_ENABLED=0
ARG GOOS=linux
ARG GOARCH=amd64

WORKDIR /go/src/k8s-host-device-plugin
COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go install -ldflags="-s -w"

FROM debian:buster-slim
COPY --from=build /go/bin/k8s-host-device-plugin /bin/k8s-host-device-plugin

CMD ["/bin/k8s-host-device-plugin"]
