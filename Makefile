# gmcm build entry #
# ================ #

VERSION_PACKAGE := gmcm/pkg/utils/version
OUTPUT_DIR := /proj/gmcm
ROOT_DIR := .

# ============================================
# Includes

include scripts/make-rules/release.mk
include scripts/make-rules/tools.mk
include scripts/make-rules/common.mk

# ============================================

.PHONY: build
build: build.binary

.PHONY: build.binary
build.binary: dldep lint npm.build
	@$(MAKE) go.build

.PHONY: build.multiarch
build.multiarch: dldep lint
	@$(MAKE) go.build.multiarch

.PHONY: build.rpm
build.rpm: lint dldep go.build
	@echo "=============> Start make rpm package"

.PHONY: lint
lint: tools.verify.golangci-lint
	@echo "=============> Run golangci to lint source code"
	@golangci-lint run -v -c $(ROOT_DIR)/.golangci.yaml $(ROOT_DIR)/...

.PHONY: npm.build
npm.build: tools.verify.statik
	@echo "=============> Run build web resource"
	@rm -fr /root/.npm/*
	@cd web && npm i -g --unsafe-perm && npm i --unsafe-perm && npm run build
	@echo "=============> Run package web resource"
	@statik -src=./web/dist

.PHONY: dldep
dldep: dep.download.ceph

.PHONY: clean
clean:
	@echo "=============> Clean all build output"
	@-rm -vrf $(OUTPUT_DIR)
