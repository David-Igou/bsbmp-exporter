FROM golang:alpine as builder
ENV GO111MODULE=on
ENV CGO_ENABLED=0
ARG TARGETOS
ARG TARGETARCH
WORKDIR $GOPATH/src/github.com/david-igou/bsbmp-exporter
COPY . .
RUN go mod download
RUN GOARCH=$TARGETARCH GOOS=$TARGETOS go build -v -o /go/bin/bsbmp-exporter github.com/david-igou/bsbmp-exporter

FROM alpine:edge
RUN apk --no-cache add ca-certificates
COPY --from=builder /go/bin/bsbmp-exporter /bin/bsbmp-exporter
ENTRYPOINT ["/bin/bsbmp-exporter"]
