
.PHONY: typescript
typescript:
	./scripts/generate-typescript.sh

.PHONY: go
go:
	./scripts/generate-go.sh

.PHONY: lint
lint:
	./scripts/lint.sh
