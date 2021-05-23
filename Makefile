
.PHONY: generate
generate:
	rm -rf gen
	mkdir gen
	mkdir gen/typescript
	./scripts/generate.sh
