.PHONY: publish install

publish:
	go run ./publish

install:
	go install ./...
