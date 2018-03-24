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
	@echo 'Building libreserver...'
	@go build -ldflags '-s -w -X main.Version='${VERSION}

clean:
	@echo 'Cleaning...'
	@go clean

install: build
	@echo installing executable file to /usr/bin/libreserver
	@sudo cp lectl /usr/bin/libreserver

uninstall: clean
	@echo removing executable file from /usr/bin/libreserver
	@sudo rm /usr/bin/libreserver
