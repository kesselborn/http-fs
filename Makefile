GOPATH=$(CURDIR)

build:
	go build

docker-image: http-fs-linux
	docker build -t kesselborn/http-fs:standalone -f Dockerfile.standalone .
	docker tag      kesselborn/http-fs:standalone  kesselborn/http-fs:latest

	docker build -t kesselborn/http-fs:alpine     -f Dockerfile .

release: docker-image
	docker tag kesselborn/http-fs:standalone kesselborn/http-fs:standalone
	docker push kesselborn/http-fs:standalone

	docker tag kesselborn/http-fs:standalone kesselborn/http-fs:latest
	docker push kesselborn/http-fs:latest

	docker tag kesselborn/http-fs:alpine kesselborn/http-fs:alpine
	docker push kesselborn/http-fs:alpine

http-fs-linux: http-fs.go
	docker run --rm -v $(CURDIR):$(CURDIR) -w $(CURDIR) golang bash -c "GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -a -installsuffix cgo -o $@ ."

test:
	./tests.sh
