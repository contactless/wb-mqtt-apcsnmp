GOM=gom
.PHONY: all prepare clean

GOPATH := $(HOME)/go
PATH := $(GOPATH)/bin:$(PATH)

DEB_TARGET_ARCH ?= armel

ifeq ($(DEB_TARGET_ARCH),armel)
GO_ENV := GOARCH=arm GOARM=5 CC_FOR_TARGET=arm-linux-gnueabi-gcc CGO_ENABLED=1
endif
ifeq ($(DEB_TARGET_ARCH),amd64)
GO_ENV := GOARCH=amd64 CC=x86_64-linux-gnu-gcc
endif
ifeq ($(DEB_TARGET_ARCH),i386)
GO_ENV := GOARCH=386 CC=i586-linux-gnu-gcc
endif

all: wb-mqtt-apcsnmp

prepare:
	go get -u github.com/mattn/gom

clean:
	rm -f wb-mqtt-apcsnmp

wb-mqtt-apcsnmp: apcsnmp_driver.go apcsnmp/*.go
	$(GO_ENV) $(GOM) install
	$(GO_ENV) $(GOM) build

install:
	mkdir -p $(DESTDIR)/usr/bin/ $(DESTDIR)/etc/init.d/
	install -m 0755 wb-mqtt-apcsnmp $(DESTDIR)/usr/bin/
	install -m 0755  initscripts/wb-mqtt-apcsnmp $(DESTDIR)/etc/init.d/wb-mqtt-apcsnmp

deb: prepare
	CC=arm-linux-gnueabi-gcc dpkg-buildpackage -b -aarmel -us -uc
