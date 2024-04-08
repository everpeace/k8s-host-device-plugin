FROM golang:1.22 as build
ARG CGO_ENABLED=0
ARG GOOS=linux
ARG GOARCH=amd64

WORKDIR /go/src/k8s-host-device-plugin
COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go install -ldflags="-s -w"

FROM gcr.io/distroless/static-debian12
COPY --from=build /go/bin/k8s-host-device-plugin /bin/k8s-host-device-plugin

CMD ["/bin/k8s-host-device-plugin"]
