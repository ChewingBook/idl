
.PHONY: all
all:
	make go && make typescript

.PHONY: typescript
typescript:
	./scripts/generate-typescript.sh
.PHONY: publish
publish:
	cp scripts/typescript/package* gen/typescript/
	cd ./gen/typescript && npm publish
	rm gen/typescript/package*
.PHONY: go
go:
	./scripts/generate-go.sh

.PHONY: lint
lint:
	./scripts/lint.sh
