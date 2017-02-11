BINARY=alfred-circleci
PACKAGE=github.com/bboughton/alfred-circleci

GOCMD=go
INSTALL_DIR = /usr/local/opt/alfred-circleci

.PHONY: build
build: 
	@$(GOCMD) build -o $(BINARY) $(PACKAGE)

.PHONY: clean
clean: 
	@$(RM) $(BINARY)

.PHONY: install
install: build
	@mkdir -p $(INSTALL_DIR)/{bin,etc}
	@install $(BINARY) $(INSTALL_DIR)/bin/$(BINARY)
	@install local.alfred-circleci.load-cache.plist $(INSTALL_DIR)/etc/
	@install load_cache.sh $(INSTALL_DIR)/bin/load_cache.sh
