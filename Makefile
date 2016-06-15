BIN = alfred-circleci
GO = go
INSTALL_DIR = /usr/local/opt/alfred-circleci

build: 
	@$(GO) build -o $(BIN)

clean: 
	@$(RM) $(BIN)

install: build
	@mkdir -p $(INSTALL_DIR)/{bin,etc}
	@install $(BIN) $(INSTALL_DIR)/bin/$(BIN)
	@install local.alfred-circleci.load-cache.plist $(INSTALL_DIR)/etc/
	@install load_cache.sh $(INSTALL_DIR)/bin/load_cache.sh

.PHONY: build clean install
