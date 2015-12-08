#
# INPUTS
# 	- IMAGE
#		- REGISTRY
#		- IMAGE
#
# TARGETS
# 	- docker-build
#		- docker-push
#

DOCKER ?= docker

.PHONY: docker-build
docker-build:
	$(DOCKER) build -t $(IMAGE) rootfs

.PHONY: docker-push
docker-push:
	$(DOCKER) tag -f $(IMAGE) $(REGISTRY)$(IMAGE)
	$(DOCKER) push $(REGISTRY)$(IMAGE)
