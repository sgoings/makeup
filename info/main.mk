#
# OUTPUTS
# 	- VERSION
# 	- PROJECT_NAME
#
# TARGETS
#		- info
#

VERSION := $(shell git describe --tags --exact-match 2>/dev/null || echo latest)
PROJECT_NAME := $(shell basename $(shell pwd))

%.makeup: FORCE
	@echo $(basename $(@F))
	@echo ---------------
	@cat $@
	@echo

.PHONY: help
help: conventions targets

.PHONY: targets
targets: $(MAKEUP_DIR)/**/*.makeup

.PHONY: conventions
conventions: $(MAKEUP_DIR)/conventions/**/*.makeup

.PHONY: info
info:
	@echo
	@echo Information
	@echo -----------
	@echo "Version:    ${VERSION}"
	@echo "Image:      ${IMAGE}"
	@echo

FORCE:
