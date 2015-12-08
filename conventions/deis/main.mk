include $(MAKEUP_DIR)/info/main.mk \
				$(MAKEUP_DIR)/docker/main.mk \
				$(MAKEUP_DIR)/split-org-registry/main.mk

.PHONY: build
build: docker-build

.PHONY: push
push: docker-push
