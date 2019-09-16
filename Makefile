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

build: clean
	@echo 'Building ${PROJECT_NAME}...'
	@go build -ldflags '-s -w -X main.version='${VERSION} '-o='${PROJECT_NAME}

clean:
	@echo 'Cleaning...'
	@go clean
	@rm -f ${PROJECT_NAME}

install: build
	@echo Installing executable file to /usr/local/bin/${PROJECT_NAME}
	@sudo cp ${PROJECT_NAME} /usr/local/bin/${PROJECT_NAME}
	@sudo cp ./lib/systemd/system/${PROJECT_NAME}.service /lib/systemd/system/${PROJECT_NAME}.service

uninstall: clean
	@echo Removing executable file from /usr/local/bin/${PROJECT_NAME}
	@sudo rm -f /usr/local/bin/${PROJECT_NAME}
	@sudo rm -f /lib/systemd/system/${PROJECT_NAME}.service
