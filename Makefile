PROJECT_NAME = 'libreserver'

# Version
VERSION = `date +%y.%m`

# If unable to grab the version, default to N/A
ifndef VERSION
    VERSION = "n/a"
endif

#
# Makefile options
#


# State the "phony" targets
.PHONY: all clean build install uninstall


all: build

build:
	@echo 'Building ${PROJECT_NAME}...'
	@go build -ldflags '-s -w -X main.Version='${VERSION}

clean:
	@echo 'Cleaning...'
	@go clean

install: build
	@echo Installing executable file to /usr/bin/${PROJECT_NAME}
	@sudo cp ${PROJECT_NAME} /usr/bin/${PROJECT_NAME}

uninstall: clean
	@echo Removing executable file from /usr/bin/${PROJECT_NAME}
	@sudo rm -f /usr/bin/${PROJECT_NAME}
