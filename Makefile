# Copyright 2025 The ChaosBlade Authors
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

.PHONY: build clean all linux_amd64 darwin_amd64 linux_arm64 darwin_arm64

# Version information management
# Priority: use environment variable BLADE_VERSION, otherwise auto-get version from Git Tag
ifneq ($(BLADE_VERSION),)
    # If environment variable BLADE_VERSION is set, use it directly
    # BLADE_VERSION is already defined in environment variables
else
    # If environment variable BLADE_VERSION is not set, try to get from Git Tag
    GIT_TAG := $(shell git describe --tags --abbrev=0 2>/dev/null || echo "")
    ifeq ($(GIT_TAG),)
        # If no Git Tag exists, use default version
        BLADE_VERSION := 1.8.0
    else
        # Extract version number from Git Tag (remove v prefix)
        BLADE_VERSION := $(shell echo $(GIT_TAG) | sed 's/^v//')
    endif
endif

# Export version number for use by other scripts
export BLADE_VERSION
# 项目信息
PROJECT_NAME := chaosblade-exec-middleware
BINARY_NAME := chaos_middleware
BUILD_DIR := target
VERSION := $(BLADE_VERSION)

# 构建环境变量
GO := go
GOOS ?= $(shell go env GOOS)
GOARCH ?= $(shell go env GOARCH)

# 获取真实的主机平台信息（不受当前环境变量影响）
HOST_GOOS := $(shell env -u GOOS go env GOOS)
HOST_GOARCH := $(shell env -u GOARCH go env GOARCH)

# 交叉编译时禁用 CGO
# 如果环境变量已设置则使用环境变量，否则根据目标平台判断
ifndef CGO_ENABLED
  ifeq ($(GOOS),$(HOST_GOOS))
    ifeq ($(GOARCH),$(HOST_GOARCH))
      CGO_ENABLED := 1
    else
      CGO_ENABLED := 0
    endif
  else
    CGO_ENABLED := 0
  endif
endif

# 构建标志
BUILD_TAGS := netgo,osusergo,disablecgo
LDFLAGS := -ldflags="-s -w -X github.com/chaosblade-io/chaosblade-exec-middleware/version.BladeVersion=$(VERSION)"

# 静态链接标志（仅在启用CGO时使用）
ifeq ($(GOOS),linux)
	# 只有在当前平台也是 Linux 且启用 CGO 时才启用外部静态链接
	ifeq ($(HOST_GOOS),linux)
	  ifeq ($(GOARCH),$(HOST_GOARCH))
	    ifeq ($(CGO_ENABLED),1)
		LDFLAGS := -ldflags="-linkmode external -extldflags -static -s -w -X github.com/chaosblade-io/chaosblade-exec-middleware/version.BladeVersion=$(VERSION)"
	    endif
	  endif
	endif
endif

# 构建目标目录
BUILD_TARGET_DIR := $(BUILD_DIR)/chaosblade-$(VERSION)-$(GOOS)_$(GOARCH)
BUILD_BIN_DIR := $(BUILD_TARGET_DIR)/bin
BUILD_YAML_DIR := $(BUILD_TARGET_DIR)/yaml

# YAML 文件
SPEC_YAML := chaosblade-middleware-spec-$(VERSION).yaml
SPEC_YAML_PATH := $(BUILD_YAML_DIR)/$(SPEC_YAML)

# 构建所有平台
build: pre_build build_yaml build_binary

# 预构建准备
pre_build:
	@echo "Building $(PROJECT_NAME) version $(VERSION) for $(GOOS)/$(GOARCH)"
	@mkdir -p $(BUILD_BIN_DIR) $(BUILD_YAML_DIR)

# 构建 YAML 规范文件（使用当前平台的 Go）
build_yaml: build/spec.go
	@echo "Generating YAML specification..."
	env GOOS= GOARCH= go run $< $(SPEC_YAML_PATH)

# 构建二进制文件
build_binary: main.go
	@echo "Building binary..."
	@echo "CGO_ENABLED=$(CGO_ENABLED), GOOS=$(GOOS), GOARCH=$(GOARCH)"
	@echo "BUILD_TAGS=$(BUILD_TAGS)"
	@CGO_ENABLED=$(CGO_ENABLED) GOOS=$(GOOS) GOARCH=$(GOARCH) $(GO) build -tags=$(BUILD_TAGS) $(LDFLAGS) -o $(BUILD_BIN_DIR)/$(BINARY_NAME) $<

# Linux AMD64 构建
linux_amd64:
	@$(MAKE) build GOOS=linux GOARCH=amd64

# Linux ARM64 构建
linux_arm64:
	@$(MAKE) build GOOS=linux GOARCH=arm64

# Darwin AMD64 构建
darwin_amd64:
	@$(MAKE) build GOOS=darwin GOARCH=amd64

# Darwin ARM64 构建
darwin_arm64:
	@$(MAKE) build GOOS=darwin GOARCH=arm64

# 构建所有支持的平台
build_all: linux_amd64 linux_arm64 darwin_amd64 darwin_arm64

# 测试
test:
	@echo "Running tests..."
	@$(GO) test -race -coverprofile=coverage.txt -covermode=atomic ./...

# 清理构建结果
clean:
	@echo "Cleaning build artifacts..."
	@$(GO) clean ./...
	@rm -rf $(BUILD_DIR)
	@rm -rf build/image/blade/chaosblade-$(VERSION)

# 显示版本信息
version:
	@echo "Version: $(VERSION)"
	@echo "Git Tag: $(shell git describe --tags --abbrev=0 2>/dev/null || echo "not available")"
	@echo "Build OS: $(GOOS)"
	@echo "Build Arch: $(GOARCH)"

.PHONY: format
format: license-format
	@echo "Running goimports and gofumpt to format Go code..."
	@./hack/update-imports.sh
	@./hack/update-gofmt.sh

.PHONY: license-check
license-check:
	@echo "Checking license headers..."
	docker run -it --rm -v $(shell pwd):/github/workspace ghcr.io/korandoru/hawkeye check

.PHONY: license-format
license-format:
	@echo "Formatting license headers..."
	docker run -it --rm -v $(shell pwd):/github/workspace ghcr.io/korandoru/hawkeye format

.PHONY: verify
verify:
	@echo "Verifying Go code formatting and import order..."
	@./hack/verify-gofmt.sh
	@./hack/verify-imports.sh

# 帮助信息
.PHONY: help
help:
	@echo "Available targets:"
	@echo "  build            - Build for current platform"
	@echo "  linux_amd64      - Build for Linux AMD64"
	@echo "  linux_arm64      - Build for Linux ARM64"
	@echo "  darwin_amd64     - Build for Darwin AMD64"
	@echo "  darwin_arm64     - Build for Darwin ARM64"
	@echo "  build_all        - Build for all supported platforms"
	@echo "  test             - Run tests"
	@echo "  clean            - Clean build artifacts"
	@echo "  version          - Show version information"
	@echo "  help             - Show this help message"
	@echo "  format           - Format Go code using goimports and gofumpt"
	@echo "  verify           - Verify Go code formatting and import order"
	@echo "  license-check    - Check license headers"
	@echo ""
	@echo "Environment variables:"
	@echo "  BLADE_VERSION    - Override version (default: Git tag or 1.7.4)"
	@echo "  GOOS            - Target OS (default: current OS)"
	@echo "  GOARCH          - Target architecture (default: current arch)"

