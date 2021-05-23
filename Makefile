
.PHONY: generate
generate:
	make lint
	rm -rf gen
	mkdir gen
	mkdir gen/typescript
	./scripts/generate.sh

.PHONY: lint
lint:
	./scripts/lint.sh
