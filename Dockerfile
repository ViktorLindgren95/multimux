FROM golang:1.18 as builder

ENV GOPROXY=""
ENV GOSUMDB=off
ENV GONOSUMDB="direct"

WORKDIR /workspace
# Copy the Go Module Manifest
COPY go.mod go.mod
#COPY go.sum go.sum
#I prefer running x incase stuff breaks, it just becomes verbose
RUN go mod download -x
# Copy the go source
COPY main.go main.go
# Build. We can change the OS here
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 GO111MODULE=on go build -a -o listener main.go
#Runner image here
FROM alpine
WORKDIR /
USER root
# Copies the built binary from the builder. Golang makes an Exe that works everywhere so we dont need a large image
COPY --from=builder /workspace/listener .
USER 1001
EXPOSE 3000
ENTRYPOINT ["/listener"]