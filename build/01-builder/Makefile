SOURCE_DIR = .
BINARY_DIR = bin
COMMAND_DIR = cmd

SOURCES = $(shell find $(SOURCE_DIR) -name "*.go")
BINARIES = $(addprefix $(BINARY_DIR)/,$(shell ls $(COMMAND_DIR)))

# Binaries target
.PHONY: binaries
binaries: $(BINARIES)

# Build binary with option to provide GOFLAGS and LDFLAGS
bin/%: goflags = $(if $(GOFLAGS),$(GOFLAGS),)
bin/%: ldflags = $(if $(LDFLAGS),-ldflags '$(LDFLAGS)',)
bin/%: cmd/% cmd/%/*.go $(SOURCES)
	CGO_ENABLED=1 go build $(goflags) $(ldflags) -o "$@" "./$<"
