# The name of the project
NAME := $(shell basename $(CURDIR))

# The architecture to build for
ARCH ?= amd64

ALL_ARCH := amd64 armhf

# If ARC is set to armhf, should set GOARCH to arm
GOARCH := $(ARCH)
ifeq ($(ARCH), armhf)
GOARCH := arm
endif

# This version-strategy uses git tags to set the version string
VERSION := $(shell git describe | grep -oE '[0-9]*\.[0-9]*\.[0-9]*-?[0-9]*')

# SHORT_VERSION does not include the commit number, used by packaging
SHORT_VERSION ?= $(shell git describe | grep -oE '[0-9]*\.[0-9]*\.[0-9]*')
SHORT_VERSION := $(if $(SHORT_VERSION),$(SHORT_VERSION),0.1.0)

# Names of various distribution related variables
DIST_DIR := $(CURDIR)/dist
DIST_NAME := $(DIST_DIR)/$(ARCH)/$(NAME)_$(SHORT_VERSION)_$(ARCH).deb
DIST_DSC := $(DIST_DIR)/source/$(NAME)_$(SHORT_VERSION).dsc
DIST_SOURCE := $(DIST_DIR)/source/$(NAME)_$(SHORT_VERSION).tar.xz

DIRS := $(addprefix bin/, $(GOARCH)) $(addprefix $(DIST_DIR)/, $(ARCH)) $(DIST_DIR)/source

# Go related variables
BIN_DIR := $(GOPATH)/bin
GODEP_BIN := $(BIN_DIR)/dep
LINTTOOL := $(BIN_DIR)/revive
GO := $(shell which go1.13.4)
GO := $(if $(GO),$(GO),/usr/lib/go-1.10/bin/go)
GOPATH ?= $(shell pwd)
PROJECT ?= $(shell $(GO) list)
PROJECT := $(if $(PROJECT),$(PROJECT),{{.Project}})

# List of binaries to build
BINS := $(shell ls $(GOPATH)/src/$(PROJECT)/cmd)

# List of packages to include in the binaries
PKGS := $(shell ls $(GOPATH)/src/$(PROJECT)/pkg)

dirs: $(DIRS)
$(DIRS):
	mkdir -p $@
	ls $@

# $(GODEP):
# 	go get -u github.com/golang/dep/cmd/dep
#
# Gopkg.toml: $(GODEP)
# 	$(GODEP_BIN) init
#
# vendor:         ## Vendor the packages using dep
# vendor: $(GODEP) Gopkg.toml Gopkg.lock
# 	@echo "No vendor dir found. Fetching dependencies now..."
# 	GOPATH=$(GOPATH):. $(GODEP_BIN) ensure

build-%:
	@$(MAKE) --no-print-directory ARCH=$* build

all-build: $(addprefix build-, $(ALL_ARCH))

build: $(addprefix bin/$(GOARCH)/, $(BINS))

$(addprefix bin/$(GOARCH)/, $(BINS)): $(shell find cmd pkg vendor -name '*.go' -print)
	@echo "building: $@"
	@GOARCH=$(GOARCH) CGO_ENABLED=1 $(GO) build -v -o $@ -installsuffix "static" $(PROJECT)/cmd/$(notdir $@)
# bin/$(ARCH)/$(BIN): build-dirs $(shell find cmd pkg -regex '.*\.\(go\)')
# 	@echo "building: $@"
# 	@GOARCH=$(ARCH) CGO_ENABLED=0 go build -o $@ -installsuffix "static" cmd/$(NAME)/main.go

deb-%:
	@$(MAKE) --no-print-directory ARCH=$* deb

all-deb: $(addprefix deb-, $(ALL_ARCH))

.PHONY: deb
deb: $(DIST_NAME) $(DIST_DSC) $(DIST_SOURCE)
$(DIST_NAME) $(DIST_DSC) $(DIST_SOURCE): DIST := xenial
$(DIST_NAME) $(DIST_DSC) $(DIST_SOURCE): $(DIRS) $(shell find cmd pkg vendor -name '*.go' -print)
	@dpkg-buildpackage -us -uc -a $(ARCH) --target-arch $(ARCH)
	@mv ../$(NAME)_$(SHORT_VERSION)_$(ARCH).deb $(DIST_NAME)
	@mv ../$(NAME)_$(SHORT_VERSION).dsc $(DIST_DSC)
	@mv ../$(NAME)_$(SHORT_VERSION).tar.xz $(DIST_SOURCE)

# .PHONY: test
# test: vendor
# 	$(GO) test -cover -timeout 30s ./pkg/...

.PHONY: lint
lint: $(LINTTOOL) vendor
	gofmt -l ./cmd ./pkg
	revive -config=revive.toml -formatter=stylish ./cmd/... ./pkg/...
	$(GO) vet ./cmd/... ./pkg/...

$(LINTTOOL):
	$(GO) get -u github.com/mgechev/revive

prefix ?= /usr/local
exec_prefix ?= $(prefix)
bindir ?= $(exec_prefix)/bin
sysconfigdir ?= /etc
.PHONY: install
install: build
	install -D -t $(DESTDIR)/$(bindir)/ $(addprefix bin/$(GOARCH)/, $(BINS))
	install -D conf/conf.toml $(DESTDIR)/$(sysconfigdir)/{{.Name}}.toml

.PHONY: clean
clean:
	rm -rf .go bin
