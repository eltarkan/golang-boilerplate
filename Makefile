API = api.yaml

.PHONY: models
models: tools/swagger
	$(call print-target)
	find ./models -type f -not -name '*_test.go' -delete
	./tools/swagger generate model --spec=docs/swagger/$(API)
