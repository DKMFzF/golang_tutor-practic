run:
	@echo "[ Start app... ]"
	@go run ./cmd/modules/main.go
	@echo "[ Finish ]"

test-all:
	@echo "[ Start all test in project... ]"
	@go test -v -count 1 ./... 
	@echo "[ Finish ]"

test-all-cache:
	@echo "[ Start all test in project (in cache)... ]"
	@go test -v ./... 
	@echo "[ Finish ]"

test-clean:
	@echo "[ Clean cache for tests - OK ]"
	@go clean -testcache

