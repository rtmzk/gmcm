# Makefile helper functions for tools #
# =================================== #

TOOLS ?= $(BLOCKER_TOOLS)
DEPS ?= $(DEPEENDENCIES)
DOWNLOAD_DIR ?= $(ROOT_DIR)/static

.PHONY: tools.install
tools.install: $(addprefix tools.install., $(TOOLS))

.PHONY: tools.install.%
tools.install.%:
	@echo "=============> Installing $*"
	@$(MAKE) install.$*

.PHONY: tools.verify.%
tools.verify.%:
	@if ! which $* &> /dev/null; then $(MAKE) tools.install.$*; fi

.PHONY: install.golangci-lint
install.golangci-lint:
	@$(GO) install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.43.0
	@golangci-lint completion bash > $(HOME)/.golangci-lint.bash
	@if ! grep -q .golangci-lint.bash $(HOME)/.bashrc; then echo "source \$$HOME/.golangci-lint.bash" >> $(HOME)/.bashrc;fi

.PHONY: install.statik
install.statik:
	@$(GO) install github.com/rakyll/statik@latest

.PHONY: dep.download
dep.download: $(addprefix dep.download.,$(DEPS))

.PHONY: dep.download.%
dep.download.%:
	@echo "=============> Downloading $*"
	@$(MAKE) download.$*

.PHONY: download.ceph
download.ceph:
	@echo "=============> Starting download ceph from $(CEPH_PACKAGE_DOWNLOADURL)"
	@curl -Lk "$(CEPH_PACKAGE_DOWNLOADURL)" -o $(DOWNLOAD_DIR)/ceph.tar.gz
	@echo "=============> End download ceph"



