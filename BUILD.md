# ChaosBlade Exec Middleware Build Guide

## Overview

This project uses Makefile for build management, supporting multi-platform compilation and automated version number management.

## Version Management

- **Automatic Version Number**: Prioritizes Git Tag major version number, defaults to `1.7.4` if none available
- **Version Format**: `1.7.4` (only contains major version number, no Git commit information)
- **Version Embedding**: Version number is embedded into binary files through compile-time link flags

## Build Targets

### Basic Build

```bash
# Build for current platform
make build

# Build and clean
make all
```

### Multi-Platform Build

```bash
# Linux AMD64
make linux_amd64

# Linux ARM64
make linux_arm64

# Darwin AMD64
make darwin_amd64

# Darwin ARM64
make darwin_arm64

# Windows AMD64
make windows_amd64

# Build all supported platforms
make build_all
```

### Special Build

```bash
# Run tests
make test

# Clean build artifacts
make clean
```

### Information Display

```bash
# Show version information
make version

# Show help information
make help
```

## Environment Variables

- `BLADE_VERSION`: Override version number (default: Git tag or 1.7.4)
- `GOOS`: Target operating system (default: current system)
- `GOARCH`: Target architecture (default: current architecture)

## Build Artifacts

Build artifacts are located in the `target/` directory with the following naming format:

```
chaosblade-{version}-{os}_{arch}/
├── bin/
│   └── chaos_middleware
└── yaml/
    └── chaosblade-middleware-spec-{version}.yaml
```

### Examples

```bash
# Directory generated after executing make linux_amd64
target/chaosblade-1.7.4-linux_amd64/
├── bin/
│   └── chaos_middleware
└── yaml/
    └── chaosblade-middleware-spec-1.7.4.yaml

# Directory generated after executing make darwin_amd64
target/chaosblade-1.7.4-darwin_amd64/
├── bin/
│   └── chaos_middleware
└── yaml/
    └── chaosblade-middleware-spec-1.7.4.yaml
```

## Build Process

1. **Pre-build Preparation**: Create necessary directory structure
2. **YAML Generation**: Generate specification files using current platform's Go
3. **Binary Build**: Build binary files using target platform's Go
4. **Version Embedding**: Embed version number into binary files through link flags

## Important Notes

- CGO is automatically disabled during cross-compilation to avoid platform compatibility issues
- Static linking is enabled for Linux platform when current platform is also Linux
- YAML generation always uses current platform's Go to ensure compatibility
- Version number only contains major version number (e.g., 1.7.4), no Git commit information

## Automated Build

### Git Tag Process

1. Create Git Tag: `git tag v1.7.5`
2. Push Tag: `git push origin v1.7.5`
3. Build: `make linux_amd64` or `make build_all`
4. Artifact directory will automatically contain version number: `chaosblade-1.7.5-linux_amd64`

### CI/CD Integration

You can use the following commands in CI/CD workflows:

```bash
# Build specific version
BLADE_VERSION=1.7.5 make linux_amd64

# Build all platforms
make build_all

# Clean and rebuild
make clean && make build_all
```

## CI/CD Workflows

This project includes GitHub Actions workflows for continuous integration and release builds.

### CI Workflow (`.github/workflows/ci.yml`)

The CI workflow runs on every push and pull request, providing:

- **Multi-platform Testing**: Tests builds on Ubuntu, macOS, and Windows
- **Multi-architecture Support**: Supports both AMD64 and ARM64 architectures
- **Go Version Matrix**: Tests with Go 1.20
- **Build Verification**: Ensures all build artifacts are properly generated
- **Cache Optimization**: Caches Go modules for faster builds

**Supported Platforms:**
- Ubuntu (Linux) - AMD64 & ARM64
- macOS (Darwin) - AMD64 & ARM64  
- Windows - AMD64 only

### Release Workflow (`.github/workflows/release.yml`)

The release workflow is triggered by:

- **Git Tags**: Automatically runs when pushing tags like `v1.7.5`
- **Manual Dispatch**: Can be manually triggered with custom version numbers

**Features:**
- **Version Management**: Automatically extracts version from Git tags
- **Release Artifacts**: Creates compressed archives for each platform
- **Artifact Upload**: Uploads build artifacts to GitHub Actions
- **Multi-platform Builds**: Builds for all supported platforms
- **Build Verification**: Comprehensive verification of build outputs

### Workflow Triggers

```bash
# Trigger CI workflow (automatic)
git push origin main
git push origin feature-branch

# Trigger release workflow
git tag v1.7.5
git push origin v1.7.5

# Manual release build (via GitHub Actions UI)
# Set version: 1.7.5
```

### CI/CD Integration Examples

```bash
# Local development workflow
make clean
make build_all
make test

# CI/CD pipeline
# 1. Code changes trigger CI workflow
# 2. CI runs tests and builds on all platforms
# 3. Create release tag
git tag v1.7.5
git push origin v1.7.5
# 4. Release workflow builds and packages artifacts
# 5. Artifacts available for download
```
