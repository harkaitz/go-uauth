PROJECT=go-uauth
VERSION=1.0.0
PREFIX=/usr/local
all:
clean:
install:
config:
	json-cfg -c google-oauth,uauth -c random-string -i ~/.config.json

## -- BLOCK:go --
all: all-go
install: install-go
clean: clean-go
deps: deps-go

build/uauth$(EXE): deps
	go build -o $@ $(GO_CONF) ./cmd/uauth

all-go:  build/uauth$(EXE)
deps-go:
	mkdir -p build
install-go:
	install -d $(DESTDIR)$(PREFIX)/bin
	cp  build/uauth$(EXE) $(DESTDIR)$(PREFIX)/bin
clean-go:
	rm -f  build/uauth$(EXE)
## -- BLOCK:go --
## -- BLOCK:license --
install: install-license
install-license: 
	mkdir -p $(DESTDIR)$(PREFIX)/share/doc/$(PROJECT)
	cp LICENSE  $(DESTDIR)$(PREFIX)/share/doc/$(PROJECT)
## -- BLOCK:license --
## -- BLOCK:man --
update: update-man
update-man:
	make-h-man update
install: install-man
install-man:
	mkdir -p $(DESTDIR)$(PREFIX)/share/man/man1
	cp ./uauth.1 $(DESTDIR)$(PREFIX)/share/man/man1
## -- BLOCK:man --
