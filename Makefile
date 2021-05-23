
.PHONY: generate
generate:
	rm -rf gen
	mkdir gen
	mkdir gen/typescript
	./scripts/generate.sh

.PHONY: lint
lint:
	protolint lint -fix protos
