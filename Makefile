
BINDIR := bin
CLIENT := $(BINDIR)/client
SERVER := $(BINDIR)/server


# Build both client and server
all: build
build: build-client build-server

build-client:
	go build -o $(CLIENT) ./cmd/client

build-server:
	go build -o $(SERVER) ./cmd/server


# Build and run executables
run-client: build-client
	$(CLIENT)

run-server: build-server
	$(SERVER)


# Delete executable binaries
clean:
	rm -rf $(BINDIR)