FROM golang:1.19 as build
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
FROM gcr.io/distroless/static-debian11
COPY --from=build /go/bin/linux_amd64/k8s-host-device-plugin /bin/k8s-host-device-plugin

CMD ["/bin/k8s-host-device-plugin"]
