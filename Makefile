.PHONY: all
all: binaries

.PHONY: test
test:
	go test -coverprofile=coverage.out ./...

include build/Makefile
include test/Makefile

.PHONY: clean
clean:
	rm -f $(BINARIES)
	docker image prune -f

.PHONY: build
build:
	docker build -f build/Dockerfile -t ghcr.io/wyvernzora/telegraf-pihole:dev .



.PHONY: copy-from-k8s
copy-from-k8s:
	./etc/pihole/download-from-k8s.sh
