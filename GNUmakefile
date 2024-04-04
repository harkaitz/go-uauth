.POSIX:
.SUFFIXES:
.PHONY: all clean install check
all:
PROJECT=go-uauth
VERSION=1.0.0
PREFIX=/usr/local

config:
	json-cfg -i google-oauth,uauth -i random-string -e ~/.config.json

## -- BLOCK:go --
build/uauth$(EXE):
	mkdir -p build
	go build -o $@ $(GO_CONF) ./cmd/uauth
all: build/uauth$(EXE)
install: all
	mkdir -p $(DESTDIR)$(PREFIX)/bin
	cp build/uauth$(EXE) $(DESTDIR)$(PREFIX)/bin
clean:
	rm -f build/uauth$(EXE)
## -- BLOCK:go --
## -- BLOCK:license --
install: install-license
install-license: 
	mkdir -p $(DESTDIR)$(PREFIX)/share/doc/$(PROJECT)
	cp LICENSE $(DESTDIR)$(PREFIX)/share/doc/$(PROJECT)
## -- BLOCK:license --
## -- BLOCK:man --
install: install-man
install-man:
	mkdir -p $(DESTDIR)$(PREFIX)/share/man/man1
	cp ./uauth.1 $(DESTDIR)$(PREFIX)/share/man/man1
## -- BLOCK:man --
