.PHONY: test
test:
	@GOOS=js GOARCH=wasm go test ./...

