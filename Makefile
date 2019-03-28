GOPATH=$(CURDIR)

build:
	go build

docker-image: *.go *.crt *.key
	docker build -t kesselborn/http-fs -f Dockerfile .

release: docker-image
	docker push kesselborn/http-fs

test:
	./tests.sh
