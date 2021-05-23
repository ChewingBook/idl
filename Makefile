
.PHONY: generate
generate:
	rm -rf gen
	mkdir gen
	mkdir gen/typescript
	./scripts/generate.sh

.PHONY: lint
lint:
	docker run --volume "$(pwd):/protos" --workdir /protos yoheimuta/protolint lint -fix .
