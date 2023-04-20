# Set the default target to compile the binary for the current operating system
.DEFAULT_GOAL = build

# Define the output directory for the compiled binary
OUTDIR = .

# Define the name of the binary
BINARY = igor

# Determine the current operating system
ifeq ($(OS),Windows_NT)
  TARGET := windows
else
  UNAME_S := $(shell uname -s)
  ifeq ($(UNAME_S),Linux)
    TARGET := linux
  endif
  ifeq ($(UNAME_S),Darwin)
    TARGET := darwin
  endif
endif

windows:
	env GOOS=windows GOARCH=amd64 go build -v -o $(OUTDIR)/$(BINARY)-windows

linux:
	env GOOS=linux GOARCH=amd64 go build -v -o $(OUTDIR)/$(BINARY)-linux

darwin:
	env GOOS=darwin GOARCH=amd64 go build -v -o $(OUTDIR)/$(BINARY)-darwin

# Define the target to build the binary for the current operating system
.PHONY: build
build: $(TARGET)

# Define the target to clean up the compiled binaries
.PHONY: clean
clean:
	go clean
	rm -rf $(OUTDIR)/$(BINARY)-windows.exe $(OUTDIR)/$(BINARY)-linux $(OUTDIR)/$(BINARY)-darwin