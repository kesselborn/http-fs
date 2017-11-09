source spec.sh

container_port() {
  echo $(docker ps --format "{{.Names }}:{{ .Ports }}" | grep http-fs-testing|sed -E 's/.*:([0-9]+)->.*/\1/g')
}

before_all() {
  make docker-image
  docker run -d --name http-fs-testing -p 5000 kesselborn/http-fs
}

after_all() {
  docker kill http-fs-testing
  docker rm   http-fs-testing
}

it_should_correctly_uploaded_files() {
  port=$(container_port)
  filename=${RANDOM}-test

  assert_match "$(curl -i -T $0 localhost:${port}/${filename})" "HTTP/1.1 200 OK"
}

it_should_correctly_serve_uploaded_files() {
  port=$(container_port)
  filename=${RANDOM}-test

  curl -T $0 localhost:${port}/${filename}

  assert "$(curl -s localhost:${port}/${filename})" "$(cat $0)"
}

it_should_correctly_delete_uploaded_files() {
  port=$(container_port)
  filename=${RANDOM}-test

  curl -T $0 localhost:${port}/${filename}

  assert_match "$(curl -i -XDELETE localhost:${port}/${filename})" "HTTP/1.1 204"
  assert_match "$(curl -i localhost:${port}/${filename})" "HTTP/1.1 404 Not Found"
}

it_should_correctly_delete_tree() {
  port=$(container_port)
  filename=foo/bar/baz/${RANDOM}-test

  curl -T $0 localhost:${port}/${filename}1
  curl -T $0 localhost:${port}/${filename}2
  curl -T $0 localhost:${port}/${filename}3
  curl -T $0 localhost:${port}/$(dirname ${filename})/foo

  assert_match "$(curl -i -XDELETE localhost:${port}/$(dirname $(dirname ${filename})))" "HTTP/1.1 204"

  assert_match "$(curl -Li localhost:${port}/${filename}1)" "HTTP/1.1 404 Not Found"
  assert_match "$(curl -Li localhost:${port}/${filename}2)" "HTTP/1.1 404 Not Found"
  assert_match "$(curl -Li localhost:${port}/${filename}3)" "HTTP/1.1 404 Not Found"
  assert_match "$(curl -Li localhost:${port}/$(dirname ${filename})/foo)" "HTTP/1.1 404 Not Found"
  assert_match "$(curl -Li localhost:${port}/$(dirname ${filename}))" "HTTP/1.1 404 Not Found"
  assert_match "$(curl -Li localhost:${port}/$(dirname $(dirname ${filename})))"  "HTTP/1.1 404 Not Found"

  assert_match "$(curl -Li localhost:${port}/$(dirname $(dirname $(dirname ${filename}))))" "HTTP/1.1 200"
}


run_tests
