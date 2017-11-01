VERSION=1.0.2

BUILDTOOL=govendor
VERSIONSTRING=$(VERSION)-`date -u +%Y%m%d.%H%M%S`
LDFLAGS="-X github.com/riussi/cryptovalues/cmd.compiled=`date -u +%Y%m%d.%H%M%S` -X github.com/riussi/cryptovalues/cmd.version=$(VERSIONSTRING)"

default: all

all: build-osx build-linux build-windows

build-osx:
	GOOS=darwin $(BUILDTOOL) build -ldflags $(LDFLAGS)
	mv cryptovalues cryptovalues-osx-$(VERSIONSTRING)

build-linux:
	GOOS=linux $(BUILDTOOL) build -ldflags $(LDFLAGS)
	mv cryptovalues cryptovalues-linux-$(VERSIONSTRING)

build-windows:
	GOOS=windows $(BUILDTOOL) build -ldflags $(LDFLAGS)
	mv cryptovalues.exe cryptovalues-win-$(VERSIONSTRING).exe
