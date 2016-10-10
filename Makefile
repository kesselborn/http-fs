GOPATH=$(CURDIR)

build:
	go build

docker-image: http-fs-linux
	docker build -t http-fs:standalone -f Dockerfile.standalone .
	docker build -t http-fs:alpine     -f Dockerfile.alpine     .

release:
	docker tag http-fs:standalone kesselborn/http-fs:standalone
	docker push kesselborn/http-fs:standalone

	docker tag http-fs:alpine kesselborn/http-fs:alpine
	docker push kesselborn/http-fs:alpine

http-fs-linux: http-fs.go
	docker run --rm -v $(CURDIR):$(CURDIR) -w $(CURDIR) golang bash -c "GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -a -installsuffix cgo -o $@ ."
