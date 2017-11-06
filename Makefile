VERSION=1.0.4

BUILDTOOL=govendor
VERSIONSTRING=$(VERSION)-`date -u +%Y%m%d.%H%M%S`
LDFLAGS="-X github.com/riussi/cryptovalues/cmd.compiled=`date -u +%Y%m%d.%H%M%S` -X github.com/riussi/cryptovalues/cmd.version=$(VERSIONSTRING)"

default: all

all: osx linux windows

osx:
	GOOS=darwin $(BUILDTOOL) build -ldflags $(LDFLAGS)
	mv cryptovalues cryptovalues-osx-$(VERSIONSTRING)

linux:
	GOOS=linux $(BUILDTOOL) build -ldflags $(LDFLAGS)
	mv cryptovalues cryptovalues-linux-$(VERSIONSTRING)

windows:
	GOOS=windows $(BUILDTOOL) build -ldflags $(LDFLAGS)
	mv cryptovalues.exe cryptovalues-win-$(VERSIONSTRING).exe
