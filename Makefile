all: build
.PHONY: all

SHELL :=/bin/bash -euEo pipefail -O inherit_errexit

MAKE_REQUIRED_MIN_VERSION:=4.2 # for SHELLSTATUS
GO_REQUIRED_MIN_VERSION ?=1.20

help:
	$(info The following make targets are available:)
	@$(MAKE) -f $(firstword $(MAKEFILE_LIST)) --print-data-base --question no-such-target 2>&1 | grep -v 'no-such-target' | \
	grep -v -e '^no-such-target' -e '^makefile' | \
	awk '/^[^.%][-A-Za-z0-9_]*:/	{ print substr($$1, 1, length($$1)-1) }' | sort -u
.PHONY: help

GO ?=go
GO_MODULE ?=$(shell $(GO) list -m)$(if $(filter $(.SHELLSTATUS),0),,$(error failed to list go module name))
GOPATH ?=$(shell $(GO) env GOPATH)
GOOS ?=$(shell $(GO) env GOOS)
GOEXE ?=$(shell $(GO) env GOEXE)
GOFMT ?=gofmt
GO_PACKAGE ?=$(shell $(GO) list -m -f '{{ .Path }}' || echo 'no_package_detected')
GO_PACKAGES ?=./...
GO_LD_FLAGS ?=
GO_BUILD_FLAGS ?=-trimpath

GO_VERSION :=$(shell $(GO) version | sed -E -e 's/.*go([0-9]+.[0-9]+.[0-9]+).*/\1/')

SWAGGER ?=$(GO) run ./vendor/github.com/go-swagger/go-swagger/cmd/swagger

SCYLLADB_SWAGGER_PATH ?=scylladb
SCYLLADB_SWAGGER_GEN_PATH ?=scylladb/gen

# We need to force locale so different envs sort files the same way for recursive traversals
diff :=LC_COLLATE=C diff --no-dereference -N

# $1 - required version
# $2 - current version
define is_equal_or_higher_version
$(strip $(filter $(2),$(firstword $(shell printf '%s\n%s' '$(1)' '$(2)' | sort -V -r -b))))
endef

# $1 - program name
# $2 - required version variable name
# $3 - current version string
define require_minimal_version
$(if $($(2)),\
$(if $(strip $(call is_equal_or_higher_version,$($(2)),$(3))),,$(error `$(1)` is required with minimal version "$($(2))", detected version "$(3)". You can override this check by using `make $(2):=`)),\
)
endef

ifneq "$(MAKE_REQUIRED_MIN_VERSION)" ""
$(call require_minimal_version,make,MAKE_REQUIRED_MIN_VERSION,$(MAKE_VERSION))
endif

ifneq "$(GO_REQUIRED_MIN_VERSION)" ""
$(call require_minimal_version,$(GO),GO_REQUIRED_MIN_VERSION,$(GO_VERSION))
endif

# $1 - application name
# $2 - spec file to use
# $4 - template directory
# $3 - target directory
define generate-swagger-client
	$(SWAGGER) generate client --name='$(1)' --spec='$(2)' --template-dir='$(3)' --target='$(4)'
endef

update-scylladbclient:
	$(call generate-swagger-client,scylladb_v1,$(SCYLLADB_SWAGGER_PATH)/scylladb_v1.json,$(SCYLLADB_SWAGGER_PATH)/templates,$(SCYLLADB_SWAGGER_GEN_PATH)/v1)
	$(call generate-swagger-client,scylladb_v2,$(SCYLLADB_SWAGGER_PATH)/scylladb_v2.json,$(SCYLLADB_SWAGGER_PATH)/templates,$(SCYLLADB_SWAGGER_GEN_PATH)/v2)
.PHONY: update-scylladbclient

update-gofmt:
	find ./ -name '*.go' -not -path './vendor/*' -exec $(GOFMT) -s -w {} '+'
.PHONY: update-gofmt

verify-generated: tmpdir:=$(shell mktemp -d)
verify-generated:
	find ./ -mindepth 1 -maxdepth 1 -not -name '\.*' -not -name 'vendor' -exec cp -a -t '$(tmpdir)' {} +
	cd '$(tmpdir)' && $(GO) mod tidy && $(GO) mod vendor && $(GO) mod verify
	+@$(MAKE) -C '$(tmpdir)' update
	$(diff) -q -r --exclude='\.*' ./  '$(tmpdir)/'
.PHONY: verify-generated

verify: verify-generated
	$(GO) vet ./...
	$(GO) mod verify
.PHONY: verify

# $1 - temporary directory
define restore-deps
	ln -s $(abspath ./) "$(1)"/current
	cp -R -H ./ "$(1)"/updated
	$(RM) -r "$(1)"/updated/vendor
	cd "$(1)"/updated && $(GO) mod tidy && $(GO) mod vendor && $(GO) mod verify
	cd "$(1)" && $(diff) -r {current,updated}/vendor/ > updated/deps.diff || true
endef

verify-deps: tmp_dir:=$(shell mktemp -d)
verify-deps:
	$(call restore-deps,$(tmp_dir))
	@echo $(diff) "$(tmp_dir)"/{current,updated}/go.mod
	@     $(diff) "$(tmp_dir)"/{current,updated}/go.mod || ( echo '`go.mod` content is incorrect - did you run `go mod tidy`?' && false )
	@echo $(diff) "$(tmp_dir)"/{current,updated}/go.sum
	@     $(diff) "$(tmp_dir)"/{current,updated}/go.sum || ( echo '`go.sum` content is incorrect - did you run `go mod tidy`?' && false )
	@echo $(diff) '$(tmp_dir)'/{current,updated}/deps.diff
	@     $(diff) '$(tmp_dir)'/{current,updated}/deps.diff || ( \
		echo "ERROR: Content of 'vendor/' directory doesn't match 'go.mod' configuration and the overrides in 'deps.diff'!" && \
		echo 'Did you run `go mod vendor`?' && \
		echo "If this is an intentional change (a carry patch) please update the 'deps.diff' using 'make update-deps-overrides'." && \
		false \
	)
.PHONY: verify-deps

update: update-gofmt update-scylladbclient
.PHONY: update

build:
	CGO_ENABLED=0 $(GO) build $(GO_BUILD_FLAGS) $(GO_LD_FLAGS) $(GO_PACKAGES)
.PHONY: build
