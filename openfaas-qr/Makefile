PROJECT?=github.com/droid42/openfaasqr-armhf
APP?=openfaasqr-armhf
PORT?=8000
RELEASE?=0.0.1
COMMIT?=$(shell git rev-parse --short HEAD)
BUILD_TIME?=$(shell date -u '+%Y-%m-%d_%H:%M:%S')

GOOS?=linux
GOARCH?=amd64

build: clean
	CGO_ENABLED=0 GOOS=${GOOS} GOARCH=${GOARCH} go build \
    		-ldflags "-s -w -X ${PROJECT}/version.Release=${RELEASE} \
    		-X ${PROJECT}/version.Commit=${COMMIT} -X ${PROJECT}/version.BuildTime=${BUILD_TIME}" \
    		-o ${APP} ${PROJECT}

clean:
	rm -f ${APP}

run: ../build
	PORT=${PORT} ./${APP}

test:
	go test -v -race ./...

container: ../build
	docker build -t $(APP):$(RELEASE) .

run: container
	docker stop $(APP):$(RELEASE) || true && docker rm $(APP):$(RELEASE) || true
	docker run --name ${APP} -p ${PORT}:${PORT} --rm \
		-e "PORT=${PORT}" \
		$(APP):$(RELEASE)

push: container
	docker push $(CONTAINER_IMAGE):$(RELEASE)