build:
	go build -o go-anything -v -ldflags "-X main.Env=dev" -tags=jsoniter

prod: 
	go build -o go-anything -v -ldflags "-X main.Env=service" -tags=jsoniter

remove:
	rm -rf go-anything

.PHONY: build remove prod
