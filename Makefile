GIT_REV?=$$(git rev-parse --short HEAD)
DATE?=$$(date -u +"%Y-%m-%dT%H:%M:%SZ")
VERSION?=$$(git describe --tags --always)
LDFLAGS="-s -w -X github.com/modulehub/mh/cmd.version=$(VERSION)-$(GIT_REV) -X main.date=$(DATE)"
goos?=$$(go env GOOS)
goarch?=$$(go env GOARCH)

GREEN  := $(shell tput -Txterm setaf 2)
YELLOW := $(shell tput -Txterm setaf 3)
CYAN   := $(shell tput -Txterm setaf 6)
RESET  := $(shell tput -Txterm sgr0)

.PHONY: default help release

default: help

## Build:
prepare: ## Download depencies and prepare dev env
	@pre-commit install
	@go mod download
	@go mod vendor

build:
	## Builds the mh binary
	set -x
	@go build -ldflags=$(LDFLAGS) -o ./bin/mh main.go

build-ci: ## Optimized build for CI
	@echo $(goos)/$(goarch)
	go build -ldflags=$(LDFLAGS) -o ./bin/mh_$(goos)_$(goarch) .
	cd ./bin && tar -czvf mh_$(goos)_$(goarch).tar.gz ./mh_$(goos)_$(goarch) ../LICENSE && cd ./..

release: ## Release with a new tag. Use like this: 'VERSION=v0.0.1 make release'
	git-chglog --next-tag $(VERSION) -o CHANGELOG.md
	git add CHANGELOG.md
	git commit -m "chore: update changelog for $(VERSION)"
	git tag $(VERSION)
	git push origin main $(VERSION)

## Test:
coverage:  ## Run test coverage suite
	@go test ./... -coverprofile=cov.out
	@go tool cover -html=cov.out
	@rm cov.out

## Help:
help: ## Show this help.
	@echo ''
	@echo 'Usage:'
	@echo '  ${YELLOW}make${RESET} ${GREEN}<target>${RESET}'
	@echo ''
	@echo 'Targets:'
	@awk 'BEGIN {FS = ":.*?## "} { \
		if (/^[a-zA-Z_-]+:.*?##.*$$/) {printf "    ${YELLOW}%-20s${GREEN}%s${RESET}\n", $$1, $$2} \
		else if (/^## .*$$/) {printf "  ${CYAN}%s${RESET}\n", substr($$1,4)} \
		}' $(MAKEFILE_LIST)

chglog: ## Generate CHANGELOG.md
	@git-chglog -o CHANGELOG.md

run:
	go run main.go $(filter-out $@,$(MAKECMDGOALS))

%:      # thanks to chakrit
	@:    # thanks to William Pursell
