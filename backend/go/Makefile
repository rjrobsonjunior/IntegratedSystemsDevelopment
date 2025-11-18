# Name of the final binary
BINARY = app

# Optional: manually specify where OpenBLAS lives
# (Uncomment if you installed OpenBLAS somewhere unusual)
# OPENBLAS_PATH = /usr/lib/x86_64-linux-gnu
# CGO_LDFLAGS = -L$(OPENBLAS_PATH) -lopenblas

# Default: Let the system linker find libopenblas
export CGO_LDFLAGS = -lopenblas

# Build command
build:
	@echo "→ Building with OpenBLAS..."
	go build -o $(BINARY) .

# Run the program
run: build
	./$(BINARY)

# Clean build artifacts
clean:
	rm -f $(BINARY)

# Show if OpenBLAS is being dynamically linked
check:
	ldd $(BINARY) | grep -i openblas || echo "OpenBLAS NOT linked!"
