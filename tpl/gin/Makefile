GOPATH:=$(shell go env GOPATH)

#git describe --tags
VERSION=0.0.1
BUILD=  `date +%FT%T%z`
COMMIT=`git rev-parse HEAD`
BRANCH=`git branch`

#LDFLAGS=-ldflags "-w -s -X main.version=${VERSION} -X main.build=${BUILD} -X main.commit=${COMMIT} -X main.branch=${BRANCH}"
LDFLAGS=-ldflags "-X main.version=${VERSION} -X main.build=${BUILD} -X main.commit=${COMMIT} -X main.branch=${BRANCH}"

DOCKER_IMAGE={{ .AppName }}
APP_NAME={{ .AppName }}


.PHONY: dev

dep:
	go get

dev:
	CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -gcflags "-N -l" ${LDFLAGS}  -o ${APP_NAME}

ini:
	go generate assets/conf/conf.go

lua:
	go generate assets/lua/lua.go

doc:
	swag init -g server/router.go

run:
	go run main.go start

cleandb:
	go run main.go clean -r

linux:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -gcflags "-N -l" ${LDFLAGS} -o ${APP_NAME}_linux

release:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build ${LDFLAGS} -o ${APP_NAME}

image:
	docker build --rm -t $(DOCKER_IMAGE) .

image-base:
	docker build --rm -t golang:base -f base.dockerfile .

image-local:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -gcflags "-N -l" ${LDFLAGS} -o ${APP_NAME}_linux
	docker build --rm -t $(DOCKER_IMAGE) -f Dockerfile_local .


exec:
	docker run -i -t --rm=true -p 9000:3000 --name="$(APP_NAME)" --entrypoint /bin/sh $(DOCKER_IMAGE)

docker-clean:
	@echo "Remove all non running containers"
	-docker rm `docker ps -q -f status=exited`
	@echo "Delete all untagged/dangling (<none>) images"
	-docker rmi `docker images -q -f dangling=true`
	@echo "Delete all volume dangling"
	-docker volume prune -f
	@echo "Delete all networks dangling"
	-docker network prune -f
