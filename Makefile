all: bin/example
test: lint unit-test

PLATFORM=local

.PHONY: bin/example
bin/example:
	@DOCKER_BUILDKIT=1 docker build . --target bin \
    --output bin/ \
    --platform ${PLATFORM}

.PHONY: unit-test
unit-test:
	@DOCKER_BUILDKIT=1 docker build . --target unit-test

.PHONY: lint
lint:
	@DOCKER_BUILDKIT=1 docker build . --target lint