VERSION=1.0.7
BUILDTOOL=go
COMPILED=`date -u +%Y%m%d.%H%M%S`
VERSIONSTRING=$(VERSION)-$(COMPILED)
LDFLAGS="-X gitlab.com/juha.ristolainen/cryptovalues/cryptovalues/cmd.compiled=$(COMPILED) -X gitlab.com/juha.ristolainen/cryptovalues/cmd.version=$(VERSIONSTRING)"

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
