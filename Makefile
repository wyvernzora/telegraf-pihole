SOURCE_DIR = .
BINARY_DIR = bin
COMMAND_DIR = cmd

SOURCES = $(shell find $(SOURCE_DIR) -name "*.go")
BINARIES = $(addprefix $(BINARY_DIR)/,$(shell ls $(COMMAND_DIR)))

.PHONY: all
all: $(BINARIES)

.PHONY: test
test:
	go test -coverprofile=coverage.out ./...

# Build binary with option to provide GOFLAGS and LDFLAGS
bin/%: goflags = $(if $(GOFLAGS),$(GOFLAGS),)
bin/%: ldflags = $(if $(LDFLAGS),-ldflags '$(LDFLAGS)',)
bin/%: cmd/% $(SOURCES)
	CGO_ENABLED=1 go build $(goflags) $(ldflags) -o "$@" "./$<"

.PHONY: clean
clean:
	rm -f $(BINARIES)
	docker image prune

.PHONY: docker
docker:
	docker build -f build/local-test/Dockerfile -t ghcr.io/wyvernzora/telegraf-pihole-local-test:dev .

.PHONY: run
run:
	docker run --rm \
		--name pylon \
		-v $(PWD)/etc/pihole:/etc/pihole:ro \
		ghcr.io/wyvernzora/telegraf-pihole-local-test:dev
