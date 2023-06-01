# go binary build entry #
# ===================== #

GO := go
SHELL := /bin/bash
BINS ?= gmcm

ifeq ($(origin VERSION),undefined)
VERSION := $(shell git describe --tags --always --match="v*")
endif
ifeq ($(origin BINVERSION),undefined)
BINVERSION := $(shell cat version)
endif

GIT_COMMIT := $(shell git rev-parse HEAD)
GO_LDFLAGS += -X $(VERSION_PACKAGE).GitVersion=$(VERSION) \
	-X $(VERSION_PACKAGE).BinVersion=$(BINVERSION) \
	-X $(VERSION_PACKAGE).GitCommit=$(GIT_COMMIT) \
	-X $(VERSION_PACKAGE).BuildDate=$(shell date +'%Y-%m-%dT%H:%M:%SZ')
GO_BUILD_FLAGS += -tags=jsoniter -ldflags "$(GO_LDFLAGS)"

PLATFORMS := linux_amd64 linux_arm64

ifeq ($(GOOS),windows)
	$(error unsupported platform. Please build binary from unix)
endif

ifeq ($(origin PLATFORM), undefined)
	ifeq ($(origin GOOS), undefined)
		GOOS := $(shell go env GOOS)
	endif

	ifeq ($(origin GOARCH), undefined)
		GOARCH := $(shell go env GOARCH)
	endif

	PLATFORM := $(GOOS)_$(GOARCH)
else
	GOOS := $(word 1, $(subst _, ,$(PLATFORM)))
	GOARCH := $(word 2, $(subst _, ,$(PLATFORM)))
endif

.PHONY: go.build.%
go.build.%:
	$(eval COMMAND := $(word 2, $(subst ., ,$*)))
	$(eval PLATFORM := $(word 1, $(subst ., ,$*)))
	$(eval OS := $(word 1, $(subst _, ,$(PLATFORM))))
	$(eval ARCH := $(word 2,$(subst _, ,$(PLATFORM))))
	@echo "=============> Build binary $(COMMAND) $(VERSION) for $(OS) $(ARCH)"
	@mkdir -p $(OUTPUT_DIR)/platforms/$(OS)/$(ARCH)
	@CGO_ENABLED=1 GOOS=$(OS) GOARCH=$(ARCH) $(GO) build $(GO_BUILD_FLAGS) -o $(OUTPUT_DIR)/platforms/$(OS)/$(ARCH)/$(COMMAND)
	@echo "=============> Remove ELF binary file's header"
	@upx $(OUTPUT_DIR)/platforms/$(OS)/$(ARCH)/$(BINS)
	@echo "=============> Packaged $(COMMAND) $(VERSION) for $(OS) $(ARCH) done. You can find them under $(OUTPUT_DIR)/platforms/$(OS)/$(ARCH)/"

.PHONY: go.build
go.build: $(addprefix go.build.,$(addprefix $(PLATFORM)., $(BINS)))

.PHONY: go.build.multiarch
go.build.multiarch: $(foreach p,$(PLATFORMS),$(addprefix go.build., $(addprefix $(p)., $(BINS))))