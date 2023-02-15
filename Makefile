.PHONY: all deps import run stop logs clean healthy

MEILISEARCH_CONTAINER_NAME := meilisearch
MEILISEARCH_CIDFILE        := $(MEILISEARCH_CONTAINER_NAME).cid

all: deps run healthy import

deps:
	go mod download

import:
	go run ./...

run: meilisearch.cid

# See: https://hub.docker.com/r/getmeili/meilisearch/
meilisearch.cid:
	docker run --rm -itd -p 7700:7700 \
		-v $(shell pwd)/meili_data:/meili_data \
		--name $(MEILISEARCH_CONTAINER_NAME) \
		--cidfile $(MEILISEARCH_CIDFILE) \
		getmeili/meilisearch:v1.0

stop:
	-docker stop $(MEILISEARCH_CONTAINER_NAME)
	rm -f $(MEILISEARCH_CIDFILE)

logs:
	docker logs -f $(MEILISEARCH_CONTAINER_NAME)

clean:
	rm -rf $(MEILISEARCH_CIDFILE) meili_data

healthy:
	@until (curl -sS localhost:7700/health | jq -r '.status' | grep -q "available") do sleep 3; done

test:
	go test ./...

.PHONY: search _req_query

Q :=

search: _req_query
	@curl -sS -X POST -H "Content-Type: application/json" 'http://localhost:7700/indexes/invoice/search' -d '{"q": "\"$(Q)\""}' | jq .

_req_query:
	$(if $(Q),,$(error need Q=))
