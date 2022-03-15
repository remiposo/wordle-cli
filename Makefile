MAKEFILE_DIR:=$(dir $(abspath $(lastword $(MAKEFILE_LIST))))
BUILD_DIR:=$(MAKEFILE_DIR)/build
CMD_DIR:=$(MAKEFILE_DIR)/cmd
VENDOR_DIR:=$(MAKEFILE_DIR)/vendor

GOOS:=$(shell go env GOOS)
GOARCH:=$(shell go env GOARCH)
TARGETS:=$(shell ls $(CMD_DIR))

.PHONY: all
all: clean vendor build

.PHONY: clean
clean:
	@rm -rf "$(BUILD_DIR)"
	@rm -rf "$(VENDOR_DIR)"

.PHONY: vendor
vendor:
	@go mod vendor

.PHONY: build
build:
	@for TARGET in $(TARGETS); do \
		cd "$(CMD_DIR)/$${TARGET}"; \
		go build -o "$(BUILD_DIR)/$(GOOS)_$(GOARCH)/$${TARGET}" -mod "vendor"; \
	done
