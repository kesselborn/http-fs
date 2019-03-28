FROM golang:alpine as builder
WORKDIR /http-fs
ENV C_ENABLED=0
COPY *.go *.crt *.key ./
RUN go build

FROM alpine
RUN apk add --no-cache ca-certificates
COPY --from=builder /http-fs/http-fs /http-fs/*.crt /http-fs/*.key /
EXPOSE 5000
CMD ["/http-fs", "-addr=0.0.0.0:5000", "-dir=/tmp"]
