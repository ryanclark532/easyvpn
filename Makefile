# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get

# Name of the main executable
MAIN_BINARY_NAME=easyvpn.exe

# Name of the separate executable
VPN_BINARY_NAME=vpnauth.exe

# Output directory
DIST_DIR=dist

all: test build

build: build-web build-main build-vpn

build-web:
	cd web && npm install && npm run build

build-main:
	$(GOBUILD) -o $(DIST_DIR)/$(MAIN_BINARY_NAME) -v

build-vpn:
	$(GOBUILD) -o $(DIST_DIR)/$(VPN_BINARY_NAME) -v ./vpn-auth

test:
	$(GOTEST) -v ./...

clean:
	$(GOCLEAN)
	rm -rf $(DIST_DIR)

run:
	$(GOBUILD) -o $(DIST_DIR)/$(MAIN_BINARY_NAME) -v ./
	./$(DIST_DIR)/$(MAIN_BINARY_NAME)

.PHONY: all build build-web build-main build-vpn test clean run

