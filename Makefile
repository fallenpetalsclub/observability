get-loki-driver:
	docker plugin install grafana/loki-docker-driver:2.8.2 --alias loki --grant-all-permissions

up:
	docker compose -f docker-compose.yaml up -d --force-recreate

down:
	docker compose -f docker-compose.yaml down

clean:
	docker compose -f docker-compose.yaml down -v

new: clean up

generate:
	go generate ./...
	rm -rf ./diagrams && \
		mkdir -p ./diagrams && \
		go run ./scripts/diagrams generate && \
		docker run --rm -ti -v ./diagrams:/diagrams -w /diagrams docker.io/zalgonoise/graphviz:ubuntu sh -c \
			'dot -Tpng observability_infra.dot > observability_infra.png'

.PHONY: dep-update
dep-update:
	go get -u ./...
	go mod tidy