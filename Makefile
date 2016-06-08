BIN = alfred-circleci
GO = go
INSTALL_DIR = /usr/local/opt/alfred-circleci

build: 
	@$(GO) build -o $(BIN)

clean: 
	@$(RM) $(BIN)

install: build
	@mkdir -p $(INSTALL_DIR)/bin
	@install $(BIN) $(INSTALL_DIR)/bin/$(BIN)

.PHONY: build clean install
