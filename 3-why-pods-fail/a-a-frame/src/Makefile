COMMIT_ID := $(shell git rev-parse HEAD)

build:
	docker build . --tag pluies/hands-on-kubernetes:a-frame-$(COMMIT_ID)

tag-current: build
	docker push pluies/hands-on-kubernetes:a-frame-$(COMMIT_ID)
install: tag-current
	docker tag pluies/hands-on-kubernetes:a-frame-$(COMMIT_ID) pluies/hands-on-kubernetes:a-frame
	docker push pluies/hands-on-kubernetes:a-frame
