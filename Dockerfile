FROM golang as builder
WORKDIR /http-fs
COPY *.go ./
RUN go build

FROM scratch
COPY --from=builder /http-fs/http-fs /http-fs
EXPOSE 5000
CMD ["/http-fs", "-addr=0.0.0.0:5000", "-dir=/tmp"]
