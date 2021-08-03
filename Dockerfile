FROM golang:1.11 as build
WORKDIR /go/src/k8s-host-device-plugin

ARG GO111MODULE=on
COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN export CGO_LDFLAGS_ALLOW='-Wl,--unresolved-symbols=ignore-in-object-files' && \
    go install -ldflags="-s -w" 

FROM debian:stretch-slim
COPY --from=build /go/bin/k8s-host-device-plugin /bin/k8s-host-device-plugin

CMD ["/bin/k8s-host-device-plugin"]
