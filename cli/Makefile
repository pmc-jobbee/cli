# SHELL := $(shell echo $$SHELL)

APP_NAME=pmc-jobbee
PROJECT_NAME := "pmc-jobbee"
VERSION := $(shell cat ../VERSION)
HASH := $(shell git describe --always)
TODAY := $(shell date "+%F %H:%M")

UNAME := $(shell uname)

ifeq ($(UNAME), Darwin)
	sed_params = -i '' -e
else ifeq ($(UNAME), Linux)
	sed_params = -i
endif

#all: build_all
all: build

clean:
	rm $(PROJECT_NAME)
	rm $(PROJECT_NAME) "$(PROJECT_NAME)_linux" "$(PROJECT_NAME)_rpi"

prepare_code: *.go
	cat app/app_version.tmpl > app/app_version.go
	sed $(sed_params) 's/__APP_NAME__/$(APP_NAME)/g' app/app_version.go
	sed $(sed_params) 's/__APP_VERSION__/$(VERSION)/g' app/app_version.go
	sed $(sed_params) 's/__APP_GIT_HASH__/$(HASH)/g' app/app_version.go
	sed $(sed_params) 's/__APP_BUILD_DATE__/$(TODAY)/g' app/app_version.go

build: prepare_code
	go build -i -v -o "$(PROJECT_NAME)"

build_linux: prepare_code
	env GOOS=linux go build -o "$(PROJECT_NAME)_linux"

build_rpi: prepare_code
	env GOOS=linux GOARCH=arm GOARM=7 go build -o "$(PROJECT_NAME)_rpi"

build_all: build build_linux build_rpi

coverage: ## Generate global code coverage report
	./tools/coverage.sh;

coverhtml: ## Generate global code coverage report in HTML
	./tools/coverage.sh html;

lint: ## Lint the files
	@golint -set_exit_status ${PKG_LIST}

msan: dep ## Run memory sanitizer
	@go test -msan -short ${PKG_LIST}

race: dep ## Run data race detector
	@go test -race -short ${PKG_LIST}

test:
	go test -v ./...