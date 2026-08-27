IMAGE_REPOSITORY ?= ghcr.io/everpeace/k8s-host-device-plugin
IMAGE_TAG        ?= $(shell git describe --tags --exact-match 2>/dev/null || git rev-parse --short HEAD)
IMAGE_TAGS       ?= $(IMAGE_TAG)

.PHONY: image-build
image-build:
	docker build $(foreach tag,$(IMAGE_TAGS),--tag=$(IMAGE_REPOSITORY):$(tag)) .

.PHONY: image-push
image-push: image-build
	$(foreach tag,$(IMAGE_TAGS),docker push $(IMAGE_REPOSITORY):$(tag);)
