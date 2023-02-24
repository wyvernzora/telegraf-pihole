.PHONY: all
all: binaries

.PHONY: test
test:
	go test -coverprofile=coverage.out ./...

include build/01-builder/Makefile

.PHONY: clean
clean:
	rm -f $(BINARIES)
	docker image prune -f

.PHONY: build
build:
	./scripts/build.sh

.PHONY: run
run:
	docker run --rm \
		--name pihole-telegraf \
		-v $(PWD)/etc/pihole:/etc/pihole:ro \
		-v $(PWD)/etc/telegraf/telegraf.conf:/etc/telegraf/telegraf.conf \
		-v $(PWD)/etc/telegraf-pihole.conf:/etc/telegraf-pihole.conf \
		ghcr.io/wyvernzora/telegraf-pihole-telegraf:dev

.PHONY: copy-from-k8s
copy-from-k8s:
	./etc/pihole/download-from-k8s.sh
