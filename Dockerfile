FROM scratch
ADD http-fs-linux /http-fs
EXPOSE 8080
CMD ["/http-fs", "-addr=0.0.0.0:8080", "-dir=/"]
