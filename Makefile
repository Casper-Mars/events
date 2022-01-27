#

API_PROTO_FILES=$(shell find ./test/api -name *.proto)

.PHONY: api
api:
	protoc --proto_path=. \
			--go_out=paths=source_relative:. \
			--go-grpc_out=paths=source_relative:. \
			$(API_PROTO_FILES)

.PHONY: run-dev
run-dev:
	docker-compose -f hack/docker-compose.yml  down
	docker-compose -f hack/docker-compose.yml  up -d

.PHONY: install-protoc-plugin
install-protoc-plugin:
	go install ./cmd/protoc-gen-go-events

.PHONY: test-protoc-plugin
test-protoc-plugin:
	protoc -I=. \
			--go_out=paths=source_relative:. \
	 		--go-events_out=paths=source_relative:. \
	 		$(API_PROTO_FILES)