BINARY := hello

clean:
	rm -rf release
.PHONY: windows
windows:
	mkdir -p release
	GOOS=windows GOARCH=amd64 go build -o release/$(BINARY)-v1.0.0-windows-amd64

.PHONY: linux
linux:
	mkdir -p release
	GOOS=linux GOARCH=amd64 go build -o release/$(BINARY)-v1.0.0-linux-amd64

.PHONY: darwin
darwin:
	mkdir -p release
	GOOS=darwin GOARCH=amd64 go build -o release/$(BINARY)-v1.0.0-darwin-amd64

.PHONY: release
release: windows linux darwin