VERSION=1.0.7
BUILDTOOL=go
COMPILED=`date -u +%Y%m%d.%H%M%S`
VERSIONSTRING=$(VERSION)-$(COMPILED)
LDFLAGS="-X github.com/riussi/cryptovalues/cmd.compiled=$(COMPILED) -X github.com/riussi/cryptovalues/cmd.version=$(VERSIONSTRING)"

default: all

all: osx linux windows

osx: 
	GOOS=darwin $(BUILDTOOL) build -ldflags $(LDFLAGS)
	mv cryptovalues-go cryptovalues-osx-$(VERSIONSTRING)

linux:
	GOOS=linux $(BUILDTOOL) build -ldflags $(LDFLAGS)
	mv cryptovalues-go cryptovalues-linux-$(VERSIONSTRING)

windows:
	GOOS=windows $(BUILDTOOL) build -ldflags $(LDFLAGS)
	mv cryptovalues-go.exe cryptovalues-win-$(VERSIONSTRING).exe
