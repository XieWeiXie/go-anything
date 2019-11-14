build:
	go build -o go-anything main.go
remove:
	rm -rf go-anything

.PHONY: build remove
