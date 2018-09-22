.DEFAULT_GOAL := build

.PHONY: build
build :
	docker-compose up -d
	docker-compose exec -T app make _build || exit 1

.PHONY: _build
_build :
	go fmt ./golang-echo/main/...
	go build ./golang-echo/main/... || exit 1

.PHONY: test
test :
	docker-compose up -d
	docker-compose exec -T app make _test || exit 1

.PHONY: _test
_test :
	go fmt ./golang-echo/test/...
	go test -cover ./golang-echo/test/... || exit 1

.PHONY: run
run :
	docker-compose up -d
	docker-compose exec -T app make _run || exit 1

.PHONY: _run
_run : _build
	./study
