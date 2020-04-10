VERSION=v0.1
T=`date +%Y-%m-%d`
BUILD="go-anything-${T}"

build:
	go build -o ${BUILD} -v -ldflags "-X main.Env=dev" -tags=jsoniter

prod: 
	go build -o ${BUILD} -v -ldflags "-X main.Env=service" -tags=jsoniter

echo:
	echo ${BUILD}

deploy:
	echo ${BUILD}
	docker run --name ${BUILD} --link mysql_for_go_anything --link redis_for_go_anything --link kafka_for_go_anything --net go-anything_go-anything-network -p 8081:8888 -d wuxiaoshen/go-anything:latest

remove:
	rm -rf ${BUILD}

.PHONY: build remove prod echo deploy
