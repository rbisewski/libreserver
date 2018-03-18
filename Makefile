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
	@echo 'Building libreoffice...'
	@go build

clean:
	@echo 'Cleaning...'
	@go clean

install: build
	@echo installing executable file to /usr/bin/libreoffice
	@sudo cp lectl /usr/bin/libreoffice

uninstall: clean
	@echo removing executable file from /usr/bin/libreoffice
	@sudo rm /usr/bin/libreoffice
