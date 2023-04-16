build:
	@go build -o bin/registryservice gradebook/cmd/registryservice
	@go build -o bin/logservice gradebook/cmd/logservice

run-registry: build
	@./bin/registryservice

run-log: build
	@./bin/logservice
