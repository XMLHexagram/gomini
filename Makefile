GO ?= go
Executable ?= gomini
TestSite ?= test-site

.PHONY: build
build:
	go build -o build/$(Executable) cmd/main.go

.PHONY: build-linux
build-linux:
	GOOS=linux GOARCH=amd64 go build -o build/$(Executable) cmd/main.go

.PHONY: clean
clean:
	rm -rf build

.PHONY: test-new-site
test-new-site: clean build
	cd build && ./$(Executable) new site $(TestSite)

.PHONY: test-genSSL-script
test-genSSL-script:  clean build test-new-site
	cd build && cp $(Executable) $(TestSite)/ && cd $(TestSite) && ./$(Executable) exec genSSL

.PHONY: copy-ssl
copy-ssl: test-new-site
	cp test/* build/$(TestSite)/

.PHONY: test-build
test-build: copy-ssl
	cd build && cp $(Executable) $(TestSite)/ && cd $(TestSite) && ./$(Executable) build

.PHONY: test-serve
test-serve: test-build
	cd build/$(TestSite) &&	git init &&./$(Executable) serve

