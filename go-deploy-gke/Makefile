GCLOUD_PROJECT_ID=sandbox-350000
DEFAULT_ARCH=amd64
PROJECT_NAME=go-sandbox

build:
	@set GOARCH=${DEFAULT_ARCH}
	@set GOOS=windows
	go build -o bin/main.exe .
	@echo "built successfully!"

run: build
	./bin/main.exe

build-image:
	docker build -t ${PROJECT_NAME} .

run-image:
	docker run -p 8080:8080 ${PROJECT_NAME}

clear-image:
	docker system prune
	docker rmi ${PROJECT_NAME}

tag-image:
	docker tag ${PROJECT_NAME} gcr.io/${GCLOUD_PROJECT_ID}/${PROJECT_NAME}

push-image:
	docker push gcr.io/${GCLOUD_PROJECT_ID}/${PROJECT_NAME}
