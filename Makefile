

KIND_CLUSTER_NAME := master
DOCKER := docker
DOCKER_IMAGE_NAME := container-builder
DOCKER_IMAGE_TAG := v1
DOCKER_BUILD_ARGS := -t $(DOCKER_IMAGE_NAME):$(DOCKER_IMAGE_TAG)

API_DEPLOYMENT := deployment/container-builder.yml

.PHONY: docker-build
docker-build:
	$(DOCKER) build $(DOCKER_BUILD_ARGS) .

.PHONY: docker-build-load
docker-build-load: docker-build
	kind load docker-image --name $(KIND_CLUSTER_NAME) $(DOCKER_IMAGE_NAME):$(DOCKER_IMAGE_TAG)

.PHONY: deploy-builder
deploy-builder:
	kubectl apply -f $(API_DEPLOYMENT)

.PHONY: deploy-secret
deploy-secret:
	kubectl apply -f secrets/podman-secret.yml