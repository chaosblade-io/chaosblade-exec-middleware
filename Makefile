.PHONY: build clean all linux_amd64 darwin_amd64 windows_amd64 linux_arm64 darwin_arm64

# 获取版本号，优先使用 Git Tag，如果没有则使用默认版本
# 只提取主要版本号，不包含提交信息
BLADE_VERSION := $(shell git describe --tags --abbrev=0 2>/dev/null | sed 's/^v//' || echo "1.7.4")

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

# Windows AMD64 构建
windows_amd64:
	@$(MAKE) build GOOS=windows GOARCH=amd64

# 构建所有支持的平台
build_all: linux_amd64 linux_arm64 darwin_amd64 darwin_arm64 windows_amd64

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

# 帮助信息
help:
	@echo "Available targets:"
	@echo "  build            - Build for current platform"
	@echo "  linux_amd64      - Build for Linux AMD64"
	@echo "  linux_arm64      - Build for Linux ARM64"
	@echo "  darwin_amd64     - Build for Darwin AMD64"
	@echo "  darwin_arm64     - Build for Darwin ARM64"
	@echo "  windows_amd64    - Build for Windows AMD64"
	@echo "  build_all        - Build for all supported platforms"
	@echo "  test             - Run tests"
	@echo "  clean            - Clean build artifacts"
	@echo "  version          - Show version information"
	@echo "  help             - Show this help message"
	@echo ""
	@echo "Environment variables:"
	@echo "  BLADE_VERSION    - Override version (default: Git tag or 1.7.4)"
	@echo "  GOOS            - Target OS (default: current OS)"
	@echo "  GOARCH          - Target architecture (default: current arch)"

