.POSIX:
.SUFFIXES:
.PHONY: all clean install check
PROJECT   =go-uauth
VERSION   =1.0.0
PREFIX    =/usr/local
BUILDDIR ?=.build

all:

## -- BLOCK:go --
.PHONY: all-go install-go clean-go $(BUILDDIR)/uauth$(EXE)
all: all-go
install: install-go
clean: clean-go
all-go: $(BUILDDIR)/uauth$(EXE)
install-go:
	mkdir -p $(DESTDIR)$(PREFIX)/bin
	cp  $(BUILDDIR)/uauth$(EXE) $(DESTDIR)$(PREFIX)/bin
clean-go:
	rm -f $(BUILDDIR)/uauth$(EXE)
##
$(BUILDDIR)/uauth$(EXE): $(GO_DEPS)
	mkdir -p $(BUILDDIR)
	go build -o $@ $(GO_CONF) ./cmd/uauth
## -- BLOCK:go --
## -- BLOCK:license --
install: install-license
install-license: README.md LICENSE
	mkdir -p $(DESTDIR)$(PREFIX)/share/doc/$(PROJECT)
	cp README.md LICENSE $(DESTDIR)$(PREFIX)/share/doc/$(PROJECT)
## -- BLOCK:license --
## -- BLOCK:man --
## -- BLOCK:man --
## -- BLOCK:sh --
install: install-sh
install-sh:
	mkdir -p $(DESTDIR)$(PREFIX)/bin
	cp bin/uauth-fake $(DESTDIR)$(PREFIX)/bin
## -- BLOCK:sh --
