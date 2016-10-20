IMAGE_VERSION=1.0.1

build-and-package: compile-linux build-image
build-deploy-dev: compile-linux build-image push-to-dev deploy-dev-image clean

compile-linux:
	CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags '-w' -o ./build/scalar

build-image:
	docker build -t ndietz/scalar .

push-to-dev:
	docker tag -f ndietz/scalar ndietz/scalar:dev
	docker push ndietz/scalar:dev

push-new-version:
	docker tag -f ndietz/scalar ndietz/scalar:$(IMAGE_VERSION)
	docker push ndietz/scalar:$(IMAGE_VERSION)

deploy-dev-image:
	kubectl create -f scalar-dev.yaml

clean:
	rm -r ./build
